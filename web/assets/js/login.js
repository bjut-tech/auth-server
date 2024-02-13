import axios from 'axios'

(function () {
  const url = new URL(location.href)
  let redirect = url.searchParams.get('redirect')

  try {
    const redirectUrl = new URL(redirect)

    const buttonCancel = document.getElementById('button_cancel')
    buttonCancel.classList.remove('is-hidden')
    buttonCancel.addEventListener('click', () => {
      location.href = redirectUrl.origin
    })
  } catch (e) {
    redirect = null
  }

  const inputUsername = document.getElementById('input_username')
  const inputPassword = document.getElementById('input_password')
  inputPassword.addEventListener('keydown', (event) => {
    if (event.key === 'Enter') {
      buttonSubmit.click()
    }
  })

  const errorMessage = document.getElementById('error-message')
  const errorMessageContainer = errorMessage.parentElement.parentElement

  const buttonSubmit = document.getElementById('button_submit')
  buttonSubmit.addEventListener('click', () => {
    const data = {
      username: inputUsername.value,
      password: inputPassword.value
    }

    errorMessageContainer.classList.add('is-hidden')
    buttonSubmit.classList.add('is-loading')

    axios.post('login', data)
      .then(() => {
        location.href = redirect ?? '/'
      })
      .catch((error) => {
        if (error.response?.data?.message) {
          errorMessage.innerText = error.response.data.message
        } else {
          errorMessage.innerText = '登录失败'
          if (error.message) {
            errorMessage.innerText += `: ${error.message}`
          }
        }
        errorMessageContainer.classList.remove('is-hidden')
        errorMessageContainer.classList.remove('is-hidden')
      })
      .finally(() => {
        buttonSubmit.classList.remove('is-loading')
      })
  })
})();
