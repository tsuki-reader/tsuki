interface Props {
    message: string | null
}

export function SuccessMessage ({ message }: Props) {
  if (message === null) {
    return <></>
  }

  return (
    <p className="text-green-500">{message}</p>
  )
}
