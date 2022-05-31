import { callApi } from './api'
import type { Post, Event, Page, Site } from './types'

const a = {
  addPost: async (body: string) => {
    await callApi<Post, void>('posts', 'POST', { id: 0, body })
    window.location.reload()
  },
  updatePost: async (post: Post) => {
    await callApi<Post, void>(`posts/${post.id}`, 'PUT', post)
    window.location.reload()
  },
  removePost: async (id: number) => {
    if (confirm('Are you sure?')) {
      await callApi<Post, void>(`posts/${id}`, 'DELETE', { id, body: '' })
      window.location.reload()
    }
  },

  addEvent: async (title: string, body: string, date: string) => {
    await callApi<Event, void>('events', 'POST', { id: 0, title, body, date })
    window.location.reload()
  },
  updateEvent: async (event: Event) => {
    await callApi<Event, void>(`events/${event.id}`, 'PUT', event)
    window.location.reload()
  },
  removeEvent: async (id: number) => {
    if (confirm('Are you sure?')) {
      await callApi<Event, void>(`events/${id}`, 'DELETE', { id, title: '', body: '', date: '' })
      window.location.reload()
    }
  },

  updateSite: async (site: Site) => {
    await callApi<Site, void>(`sites/${site.id}`, 'PUT', site)
    window.location.reload()
  },

  addPage: async (title: string, body: string) => {
    await callApi<Page, void>('pages', 'POST', { id: 0, title, body })
    window.location.reload()
  },
  updatePage: async (page: Page) => {
    await callApi<Page, void>(`pages/${page.id}`, 'PUT', page)
    window.location.reload()
  },
  removePage: async (id: number) => {
    if (confirm('Are you sure?')) {
      await callApi<Page, void>(`pages/${id}`, 'DELETE', { id, title: '', body: '' })
      window.location.href = '/'
    }
  },
}

export default a
