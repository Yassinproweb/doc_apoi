// code for search bar filter using JSON data.
fetch('./search.json')
  .then(response => response.json())
  .then(data => {
    const dataList = document.getElementById('searchList');
    const inputSearch = document.getElementById('search');
    const highlight = document.getElementById('highlight');

    data.forEach(item => {
      const option = document.createElement('option');
      option.value = item.doctor + " - " + item.specialty;
      dataList.appendChild(option);
    });

    inputSearch.addEventListener('input', function() {
      const inputValue = inputSearch.value.toLowerCase();
      let highlightText = '';

      data.forEach(item => {
        const optionText = (item.doctor + " - " + item.specialty).toLowerCase();
        if (optionText.startsWith(inputValue)) {
          highlightText = item.doctor.substring(0, inputValue.length);
        }
      });

      highlight.innerText = highlightText;
    });

    inputSearch.addEventListener('keydown', function(event) {
      if (event.key === "Tab" && dataList.firstChild) {
        inputSearch.value = dataList.firstChild.value;
      }
    });
  })
  .catch(error => console.error('Error fetching data:', error));
