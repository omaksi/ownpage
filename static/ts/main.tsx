import { initAdmin } from './admin'

const isLoggedIn = document.body.getAttribute('data-logged-in') !== ''

console.log('isLoggedIn', document.body.getAttribute('data-logged-in'))
console.log('isLoggedIn', isLoggedIn)

if (isLoggedIn) {
  initAdmin()
}
