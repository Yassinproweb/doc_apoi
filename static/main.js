const header = document.querySelector("header");

window.onscroll = function() {
  if (window.scrollY > 75) {
    header.classList.add("scroll");
  } else {
    header.classList.remove("scroll");
  }
};

// registration forms for doctors & patients
const formJoin = document.getElementById("form-join"),
  formDoc = document.getElementById("register_doctor"),
  formPat = document.getElementById("register_patient"),
  loginDoc = document.getElementById("login_doctor"),
  loginPat = document.getElementById("login_patient"),
  closeForm = document.getElementById("close-form"),
  regDoc = document.getElementById("register-doc"),
  regPat = document.getElementById("register-pat"),
  signDoc = document.getElementById("login-doc"),
  signPat = document.getElementById("login-pat");

function addHidden(p1) {
  if (!p1.classList.contains("hidden")) {
    p1.classList.add("hidden");
  };
};

regDoc.addEventListener("click", () => {
  formJoin.classList.remove("hidden");
  formDoc.classList.remove("hidden");
});

regPat.addEventListener("click", () => {
  formJoin.classList.remove("hidden");
  formDoc.classList.remove("hidden");
});

signDoc.addEventListener("click", () => {
  loginDoc.classList.remove("hidden");
  addHidden(formDoc);
});

signPat.addEventListener("click", () => {
  loginPat.classList.remove("hidden");
  addHidden(formPat);
});

closeForm.addEventListener("click", () => {
  formJoin.classList.add("hidden");
  addHidden(formDoc);
  addHidden(formPat);
  addHidden(loginDoc);
  addHidden(loginPat);
});
