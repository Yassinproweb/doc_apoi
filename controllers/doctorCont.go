package controllers

import (
	"fmt"
	"os"
	"path/filepath"
	"strings"

	"github.com/Yassinproweb/doc_apoi/models"
	"github.com/Yassinproweb/doc_apoi/utils"
	"github.com/gofiber/fiber/v2"
	"github.com/google/uuid"
)

// Fetch all doctors and render homepage
func GetDoctorsController() fiber.Handler {
	return func(c *fiber.Ctx) error {
		doctors, err := models.GetAllDoctors()
		if err != nil {
			return c.Status(fiber.StatusInternalServerError).SendString("Failed to fetch doctors")
		}

		return c.Render("dashboard", fiber.Map{
			"Doctors": doctors,
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

			const maxFileSize = 2 * 1024 * 1024 // 2MB
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

		err = models.AddDoctor(name, email, password, skill, title, location, avatar)
		if err != nil {
			return c.Status(fiber.StatusBadRequest).SendString("Registration failed: " + err.Error())
		}

		// Set only email cookie
		c.Cookie(&fiber.Cookie{
			Name:     "doctor_email",
			Value:    email,
			Path:     "/",
			HTTPOnly: true,  // prevent JS access
			SameSite: "Lax", // protect against CSRF
			Secure:   false, // only send over HTTPS (use false in localhost)
			MaxAge:   60 * 60 * 24 * 7,
		})

		nameURL := strings.ReplaceAll(strings.ToLower(name), " ", "_")
		redirectURL := fmt.Sprintf("/doctors/%s", nameURL)

		c.Set("HX-Redirect", redirectURL)
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
			MaxAge:   60 * 60 * 24 * 7,
		})

		nameURL := strings.ReplaceAll(strings.ToLower(d.Name), " ", "_")
		redirectURL := fmt.Sprintf("/doctors/%s", nameURL)

		c.Set("HX-Redirect", redirectURL)
		return c.SendStatus(fiber.StatusOK)
	}
}

// Doctor profile redirect
func DoctorRedirect() fiber.Handler {
	return func(c *fiber.Ctx) error {
		name := c.Params("name")

		d, err := models.GetDoctorByName(name)
		if err != nil {
			return c.Status(fiber.StatusNotFound).SendString("Doctor not found")
		}

		return c.Render("dr-profile", fiber.Map{
			"Doctor": d,
		})
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

			const maxFileSize = 2 * 1024 * 1024 // 2MB
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
		c.Set("HX-Redirect", "/dashboard")
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
		c.Set("HX-Redirect", "/dashboard")
		return c.SendStatus(fiber.StatusOK)
	}
}
