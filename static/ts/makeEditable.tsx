import { h, render, Fragment } from '../preact/preact.mjs'
import AdminControls from './cmp/AdminControls.js'
import SubmitCancel from './cmp/SubmitCancel.js'
import ne from './newElement.js'

function makeEditable(
  editableElements: { [key: string]: HTMLElement },
  controlRoot: HTMLElement,
  id: number,
  onRemove: (id: number) => void,
  onSubmit: (values: any) => void
  // onImage: () => void | null
) {
  function submit() {
    console.log('submit')
    const values = {}

    for (const key in editableElements) {
      const el = editableElements[key]
      values[key] = el.textContent
    }

    values['id'] = id
    console.log(values)

    onSubmit(values)
  }

  function remove() {
    onRemove(id)
  }

  // function upload() {}

  function handleToggleEdit() {
    for (const key in editableElements) {
      const el = editableElements[key]
      if (el.isContentEditable) {
        el.classList.remove('editable')
        el.contentEditable = 'false'
        el.blur()
        render(
          <AdminControls toggleEdit={handleToggleEdit} hasRemove onRemove={remove} />,
          controls
        )
      } else {
        el.classList.add('editable')
        el.contentEditable = 'true'
        el.focus()
        render(<SubmitCancel onSubmit={submit} onCancel={handleToggleEdit} />, controls)
      }
    }
  }

  const controls = ne('div', { class: 'controls' }, controlRoot, null)
  // const uploadField = ne('input', { type: 'file', name: 'image', class: 'hidden' }, controls, null)

  // uploadField.addEventListener('change', function (e) {
  //   console.log('uploadField.onchange')
  //   const target = e.target as HTMLInputElement
  //   const file = target.files[0]
  //   console.log(file)
  // })

  render(<AdminControls toggleEdit={handleToggleEdit} hasRemove onRemove={remove} />, controls)
}

export default makeEditable
