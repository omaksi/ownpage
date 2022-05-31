import { h, Fragment } from '../../preact/preact.mjs'

type AddPostProps = {}

function AddPost(props) {
  return (
    <Fragment>
      <form action="/api/posts" method="post" id="post-add-form">
        <input class="input" type="hidden" name="Id" value="0" />
        <input
          class="input"
          type="text"
          name="Body"
          placeholder="Start writing your new post here..."
        />
        <button class="button" type="submit" value="post">
          Post
        </button>
      </form>
    </Fragment>
  )
}

export default AddPost
