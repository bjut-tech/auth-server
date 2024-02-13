import axios from 'axios'

(function () {
  const url = new URL(location.href)
  if (url.searchParams.has('unauthorized')) {
    document.getElementById('container_main').classList.add('is-hidden')
    document.getElementById('container_unauthorized').classList.remove('is-hidden')
  }

  axios.get('userinfo')
    .then((response) => {
      document.querySelectorAll('.is-logged-in').forEach((el) => {
        el.classList.remove('is-hidden')
      })
      document.getElementById('claims').innerText = JSON.stringify(response.data, null, 2)
    })
    .catch((error) => {
      if (error.response?.status === 401) {
        document.querySelectorAll('.is-not-logged-in').forEach((el) => {
          el.classList.remove('is-hidden')
        })
      } else {
        console.error(error)
      }
    })

  document.getElementById('button_login').addEventListener('click', () => {
    location.href = 'login.html'
  })

  document.getElementById('button_logout').addEventListener('click', () => {
    axios.post('logout')
      .then(() => {
        if (location.search) {
          location.href = location.pathname
        } else {
          location.reload()
        }
      })
      .catch((error) => {
        console.error(error)
      })
  })
})();
