import { h, Fragment } from '../../preact/preact.mjs'

type AddEventProps = {}

function AddEvent(props) {
  return (
    <Fragment>
      <form action="/api/events" method="post" id="event-add-form">
        <input class="input" type="hidden" name="Id" value="0" />
        <input class="input" type="text" name="Title" placeholder="Event title" />
        <input class="input" type="text" name="Body" placeholder="Event description..." />
        <input class="input" type="date" name="Date" placeholder="Event date..." />

        <button class="button" type="submit" value="event">
          Post Event
        </button>
      </form>
    </Fragment>
  )
}

export default AddEvent
