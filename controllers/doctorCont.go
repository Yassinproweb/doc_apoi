package controllers

import (
	"log"
	"os"
	"path/filepath"

	"github.com/Yassinproweb/doc_apoi/models"
	"github.com/Yassinproweb/doc_apoi/utils"
	"github.com/gofiber/fiber/v2"
	"github.com/google/uuid"
)

// Doctor dashboard controller
func DoctorDashboardController() fiber.Handler {
	return func(c *fiber.Ctx) error {
		paramName := c.Params("name")
		email := c.Cookies("doctor_email")

		if email == "" {
			return c.Redirect("/doctors?mode=login")
		}

		d, err := models.GetDoctor(email)
		if err != nil {
			return c.Status(fiber.StatusUnauthorized).SendString("Invalid user or not registered")
		}

		eslug := utils.NormalizeName(d.Name)
		if eslug != paramName {
			c.Set("HX-Redirect", utils.URLer("doctors", eslug))
			return c.SendStatus(fiber.StatusOK)
		}
		return c.Render("doctors", fiber.Map{
			"Doctor": d,
		})
	}
}

// Doctor registration
func RegisterDoctorController() fiber.Handler {
	return func(c *fiber.Ctx) error {
		name := c.FormValue("name")
		email := c.FormValue("email")
		password := c.FormValue("password")
		skill := c.FormValue("skill")
		title := c.FormValue("title")
		location := c.FormValue("location")

		file, err := c.FormFile("avatar")
		avatar := "/static/imgs/pngs/doc_apoi.png" // default logo

		allowedFiles := []string{"image/jpeg", "image/png", "image/webp", "image/gif"}
		avatarsDir := "./static/avatars/"
		os.MkdirAll(avatarsDir, os.ModePerm)

		if err == nil && file != nil {
			contentType := file.Header.Get("Content-Type")
			if !utils.IsAllowedFileType(contentType, allowedFiles) {
				return c.Status(fiber.StatusBadRequest).SendString("Invalid image file! Only .jpg, .jpeg, .png, .webp and .gif are allowed.")
			}

			const maxFileSize = 2_097_152 // 2MB
			if file.Size > maxFileSize {
				return c.Status(fiber.StatusBadRequest).SendString("File too large. Max size is 2MB.")
			}

			ext := filepath.Ext(file.Filename)
			uniqueName := uuid.New().String() + ext

			if err := c.SaveFile(file, avatarsDir+uniqueName); err != nil {
				return c.Status(fiber.StatusInternalServerError).SendString("Failed to save image.")
			}

			avatar = "/static/avatars/" + uniqueName
		}

		log.Println(avatar)

		err = models.AddDoctor(name, email, password, skill, title, location, avatar)
		if err != nil {
			return c.Status(fiber.StatusInternalServerError).SendString("Registration failed: " + err.Error())
		}

		// Set only email cookie
		c.Cookie(&fiber.Cookie{
			Name:     "doctor_email",
			Value:    email,
			Path:     "/",
			HTTPOnly: true,  // prevent JS access
			SameSite: "Lax", // protect against CSRF
			Secure:   false, // only send over HTTPS (use false in localhost)
			MaxAge:   604800,
		})

		c.Set("HX-Redirect", utils.URLer("doctors", utils.NormalizeName(name)))
		return c.SendStatus(fiber.StatusCreated)
	}
}

// Doctor login
func LoginDoctorController() fiber.Handler {
	return func(c *fiber.Ctx) error {
		email := c.FormValue("email")
		password := c.FormValue("password")

		d, err := models.GetDoctor(email)
		if err != nil || !d.CheckPassword(password) {
			return c.SendStatus(fiber.StatusUnauthorized)
		}

		// Set only email cookie
		c.Cookie(&fiber.Cookie{
			Name:     "doctor_email",
			Value:    email,
			Path:     "/",
			HTTPOnly: true,  // prevent JS access
			SameSite: "Lax", // protect against CSRF
			Secure:   false, // only send over HTTPS (use false in localhost)
			MaxAge:   604800,
		})

		c.Set("HX-Redirect", utils.URLer("doctors", utils.NormalizeName(d.Name)))
		return c.SendStatus(fiber.StatusOK)
	}
}

// Update doctor profile (info + optional new avatar)
func UpdateDoctorController() fiber.Handler {
	return func(c *fiber.Ctx) error {
		email := c.Cookies("doctor_email")
		if email == "" {
			return c.SendStatus(fiber.StatusUnauthorized)
		}

		name := c.FormValue("name")
		skill := c.FormValue("skill")
		title := c.FormValue("title")
		location := c.FormValue("location")
		password := c.FormValue("password")

		// Handle optional avatar update
		file, err := c.FormFile("avatar")
		if err == nil && file != nil {
			allowedFiles := []string{"image/jpeg", "image/png", "image/webp", "image/gif"}
			contentType := file.Header.Get("Content-Type")

			if !utils.IsAllowedFileType(contentType, allowedFiles) {
				return c.Status(fiber.StatusBadRequest).SendString("Invalid image file!")
			}

			const maxFileSize = 2_097_152 // 2MB
			if file.Size > maxFileSize {
				return c.Status(fiber.StatusBadRequest).SendString("File too large. Max 2MB.")
			}

			os.MkdirAll("./static/avatars/", os.ModePerm)
			ext := filepath.Ext(file.Filename)
			uniqueName := uuid.New().String() + ext
			savePath := "./static/avatars/" + uniqueName

			if err := c.SaveFile(file, savePath); err != nil {
				return c.Status(fiber.StatusInternalServerError).SendString("Failed to save image")
			}

			avatarPath := "/static/avatars/" + uniqueName
			if err := models.UpdateDoctorAvatar(email, avatarPath); err != nil {
				return c.Status(fiber.StatusInternalServerError).SendString("Failed to update avatar")
			}
		}

		_, err = models.EditDoctor(email, name, password, skill, title, location)
		if err != nil {
			return c.SendStatus(fiber.StatusInternalServerError)
		}

		// Redirect for HTMX
		c.Set("HX-Redirect", utils.URLer("doctors", utils.NormalizeName(name)))
		return c.SendStatus(fiber.StatusOK)
	}
}

// Get edit doctor form
func EditDoctorFormController() fiber.Handler {
	return func(c *fiber.Ctx) error {
		email := c.Cookies("doctor_email")
		if email == "" {
			return c.SendStatus(fiber.StatusUnauthorized)
		}

		var d models.Doctor
		err := models.GetDoctorByEmail(email, &d)
		if err != nil {
			return c.SendStatus(fiber.StatusNotFound)
		}

		return c.Render("update-doctor", fiber.Map{"Doctor": d})
	}
}

// Doctor logout
func LogoutDoctorController() fiber.Handler {
	return func(c *fiber.Ctx) error {
		c.Cookie(&fiber.Cookie{
			Name:     "doctor_email",
			Value:    "",
			Path:     "/",
			HTTPOnly: true,  // prevent JS access
			SameSite: "Lax", // protect against CSRF
			Secure:   false, // only send over HTTPS (use false in localhost)
			MaxAge:   -1,
		})

		// Redirect for HTMX
		c.Set("HX-Redirect", "/doctors")
		return c.SendStatus(fiber.StatusOK)
	}
}
