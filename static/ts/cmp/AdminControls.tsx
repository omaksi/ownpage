import { h, Fragment } from '../../preact/preact.mjs'

type AdminControlsProps = {
  toggleEdit: () => void
  hasRemove: boolean
  onRemove?: (id: number) => void
  hasImage?: boolean
  onImage?: (id: number) => void
}

function AdminControls(props: AdminControlsProps) {
  return (
    <Fragment>
      <a href="#" onClick={props.toggleEdit}>
        <i class="fa fa-pencil" aria-hidden="true"></i>
      </a>
      {props.hasRemove && (
        <a href="#" onClick={props.onRemove}>
          <i class="fa fa-trash" aria-hidden="true"></i>
        </a>
      )}
      {props.hasImage && (
        <a href="#" onClick={props.onImage}>
          <i class="fa fa-trash" aria-hidden="true"></i>
        </a>
      )}
    </Fragment>
  )
}

export default AdminControls
