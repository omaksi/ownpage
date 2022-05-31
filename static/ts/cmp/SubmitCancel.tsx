import { h, Fragment } from '../../preact/preact.mjs'

type SubmitCancelProps = {
  onSubmit: () => void
  onCancel: () => void
}

function SubmitCancel(props: SubmitCancelProps) {
  return (
    <Fragment>
      <a href="#" onClick={props.onSubmit}>
        <i class="fa fa-check-circle-o" aria-hidden="true"></i>
      </a>
      <a href="#" onClick={props.onCancel}>
        <i class="fa fa-times-circle-o" aria-hidden="true"></i>
      </a>
    </Fragment>
  )
}

export default SubmitCancel
