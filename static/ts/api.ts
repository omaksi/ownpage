async function callApi<TBody, TResponse>(endpoint: string, method: string, body: TBody) {
  try {
    const res = await fetch(`/api/${endpoint}`, {
      method,
      body: JSON.stringify(body),
      headers: {
        'Content-Type': 'application/json',
      },
    })

    if (!res.ok) {
      throw new Error(res.statusText)
    }

    const json = await res.json()

    return json as TResponse
  } catch (err) {
    throw new Error(err)
  }
}

export { callApi }
