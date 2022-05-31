import { h, Fragment } from '../../preact/preact.mjs'

type AddPageButtonProps = {
  onClick: () => void
}

function AddPageButton(props: AddPageButtonProps) {
  return (
    <a href="#" onClick={props.onClick}>
      <i class="fa fa-plus"></i>
    </a>
  )
}

export default AddPageButton
