import { h, render, Fragment } from '../preact/preact.mjs'
import a from './actions'
import AddEvent from './cmp/AddEvent.js'
import AddPage from './cmp/AddPage.js'
import AddPageButton from './cmp/AddPageButton.js'
import AddPost from './cmp/AddPost'
import makeEditable from './makeEditable'
import ne from './newElement.js'

export function initAdmin() {
  // console.log('initAdmin')
  makePostsEditable()
  makeEventsEditable()
  makeHeaderEditable()
  renderAddPostForm()
  renderAddEventForm()
  makePageEditable()
  // renderAddPageForm()
  renderAddPageButton()
}

function makePostsEditable() {
  const posts = document.querySelectorAll<HTMLDivElement>('.post')
  if (!posts.length) {
    return
  }

  posts.forEach((post) => {
    const body = post.querySelector<HTMLDivElement>('.body')
    const id = parseInt(post.getAttribute('data-id'), 10)
    makeEditable({ body: body }, post, id, a.removePost, a.updatePost)
  })
}

function makeEventsEditable() {
  const events = document.querySelectorAll<HTMLDivElement>('.event')
  if (!events.length) {
    return
  }

  events.forEach((event) => {
    const title = event.querySelector<HTMLDivElement>('.title')
    const body = event.querySelector<HTMLDivElement>('.body')
    const id = parseInt(event.getAttribute('data-id'), 10)
    makeEditable({ body, title }, event, id, a.removeEvent, a.updateEvent)
  })
}

function makeHeaderEditable() {
  const header = document.querySelector<HTMLElement>('header')
  if (!header) {
    return
  }

  const title = header.querySelector<HTMLHeadingElement>('h1')
  const description = header.querySelector<HTMLParagraphElement>('p')
  const siteId = parseInt(header.getAttribute('data-id'), 10)
  makeEditable({ title, description }, header, siteId, null, a.updateSite)
}

function makePageEditable() {
  const page = document.querySelector<HTMLDivElement>('.page')
  if (!page) {
    return
  }

  const title = page.querySelector<HTMLHeadingElement>('h2')
  const body = page.querySelector<HTMLParagraphElement>('p')
  const id = parseInt(page.getAttribute('data-id'), 10)
  makeEditable({ title, body }, page, id, a.removePage, a.updatePage)
}

function renderAddPostForm() {
  const posts = document.querySelector<HTMLDivElement>('.posts')
  if (!posts) {
    return
  }
  const addPostWrapper = ne('div', { class: 'post-add card' }, posts, null, { prepend: true })
  render(<AddPost />, addPostWrapper)
}

function renderAddEventForm() {
  const events = document.querySelector<HTMLDivElement>('.events')
  console.log('events', events)
  if (!events) {
    return
  }
  const addEventWrapper = ne('div', { class: 'event-add card' }, events, null, {
    prepend: true,
  })
  render(<AddEvent />, addEventWrapper)
}

function renderAddPageForm() {
  const main = document.querySelector<HTMLElement>('main')
  if (!main) {
    return
  }
  main.innerHTML = ''
  render(<AddPage />, main)
}

function renderAddPageButton() {
  const menu = document.querySelector<HTMLElement>('.main-menu')
  if (!menu) {
    return
  }

  const addPageButton = ne('div', { class: 'add-page-button' }, menu, null)
  render(<AddPageButton onClick={renderAddPageForm} />, addPageButton)
}
