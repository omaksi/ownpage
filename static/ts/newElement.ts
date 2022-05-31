function ne<T extends keyof HTMLElementTagNameMap>(
  type: T,
  attributes?: { [key: string]: string },
  parentNode?: HTMLElement,
  children?: HTMLElement[],
  options?: any
): HTMLElementTagNameMap[T] {
  const el = document.createElement<T>(type)

  if (attributes) {
    for (const key in attributes) {
      el.setAttribute(key, attributes[key])
    }
  }

  if (parentNode) {
    if (options && options.prepend === true) {
      parentNode.prepend(el)
    } else {
      parentNode.appendChild(el)
    }
  }
  if (children) {
    children.forEach((child) => {
      el.appendChild(child)
    })
  }
  return el
}

export default ne
