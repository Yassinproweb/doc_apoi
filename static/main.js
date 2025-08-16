const header = document.querySelector("header");

window.onscroll = function() {
  if (window.scrollY > 75) {
    header.classList.add("scroll");
  } else {
    header.classList.remove("scroll");
  }
};


// // code for search bar filter using JSON data.
// fetch('./search.json')
//   .then(response => response.json())
//   .then(data => {
//     const dataList = document.getElementById('searchList');
//     const inputSearch = document.getElementById('search');
//     const highlight = document.getElementById('highlight');
//
//     data.forEach(item => {
//       const option = document.createElement('option');
//       option.value = item.doctor + " - " + item.specialty;
//       dataList.appendChild(option);
//     });
//
//     inputSearch.addEventListener('input', function() {
//       const inputValue = inputSearch.value.toLowerCase();
//       let highlightText = '';
//
//       data.forEach(item => {
//         const optionText = (item.doctor + " - " + item.specialty).toLowerCase();
//         if (optionText.startsWith(inputValue)) {
//           highlightText = item.doctor.substring(0, inputValue.length);
//         }
//       });
//
//       highlight.innerText = highlightText;
//     });
//
//     inputSearch.addEventListener('keydown', function(event) {
//       if (event.key === "Tab" && dataList.firstChild) {
//         inputSearch.value = dataList.firstChild.value;
//       }
//     });
//   })
//   .catch(error => console.error('Error fetching data:', error));

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
