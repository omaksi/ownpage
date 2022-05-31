import { h, Fragment } from '../../preact/preact.mjs'

type AddPageProps = {}

function AddPage(props) {
  return (
    <div class="page-add">
      <h2>Add page</h2>
      <form action="/api/pages" method="post" id="post-add-form">
        <input class="input" type="hidden" name="Id" value="0" />
        <input class="input" type="text" name="Slug" placeholder="slug, without the leading /" />
        <input class="input" type="text" name="Title" placeholder="Page title" />
        <textarea class="input" type="text" name="Body" placeholder="Page text..." />
        <button class="button" type="submit" value="post">
          Add Page
        </button>
      </form>
    </div>
  )
}

export default AddPage
