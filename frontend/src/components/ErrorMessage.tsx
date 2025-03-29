interface Props {
    message: string | null
}

export function ErrorMessage ({ message }: Props) {
  if (message === null) {
    return <></>
  }

  return (
        <p className="px-1 rounded bg-red-500 text-foreground">{message}</p>
  )
}
