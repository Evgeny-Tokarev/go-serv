const form = document.querySelector("#form")
form.addEventListener("submit", submitHandler)

function submitHandler(event) {
  event.preventDefault()
  const params = { name: event.target[0].value, address: event.target[1].value }
  const xhr = new XMLHttpRequest()
  const url = "http://localhost:9000/form"
  xhr.open("POST", url, true)

  xhr.setRequestHeader("Content-Type", "application/json")

  xhr.onreadystatechange = function () {
    if (xhr.readyState == 4 && xhr.status == 200) {
      console.log("Successfully: ", xhr.responseText)
    }
  }
  xhr.send(JSON.stringify(params))
}
