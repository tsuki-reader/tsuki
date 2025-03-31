import { ChangeEvent, useContext, useState } from 'react'
import { Modal } from '../Modal'
import { ErrorMessage } from '../ErrorMessage'
import { extensions } from '../../../wailsjs/go/models'
import { RepositoriesCreate } from '../../../wailsjs/go/backend/App'

interface Props {
  opened: boolean
  onRepoInstalled: (repos: extensions.Repository[]) => void
  onClose: React.ReactEventHandler<HTMLDialogElement>
}

export function InstallRepository ({ opened, onRepoInstalled, onClose }: Props) {
  const [valid, setValid] = useState<boolean>(false)
  const [value, setValue] = useState<string>('')
  const [errorMessage, setErrorMessage] = useState<string | null>(null)

  const isValidUrl = (url: string) => {
    try {
      /* eslint-disable no-new */
      new URL(url)
      return true
    } catch (err) {
      return false
    }
  }

  const onChange = (event: ChangeEvent<HTMLInputElement>) => {
    const newValue = event.target.value
    setValue(newValue)
    if (isValidUrl(newValue)) {
      setValid(true)
    } else {
      setValid(false)
    }
  }

  const handleResponse = (repos: extensions.Repository[]) => {
    onRepoInstalled(repos)
    setValue('')
    setErrorMessage(null)
  }

  const closeModal = (
    event: React.SyntheticEvent<HTMLDialogElement, Event>
  ) => {
    setErrorMessage(null)
    onClose(event)
  }

  const install = () => {
    RepositoriesCreate(value).then(handleResponse).catch((err) => setErrorMessage(err))
  }

  return (
    <Modal opened={opened} onClose={closeModal}>
      <div className="flex flex-col gap-4 text-center">
        <h1 className="font-bold text-3xl">Install from url</h1>
        <p>
          Install a repository by providing the URL to the remote repository
          file.
        </p>
        <input
          value={value}
          onChange={onChange}
          placeholder="https://example.com/repo.json"
          className="rounded p-2 text-foreground bg-foreground/25 border-2 border-foreground placeholder:text-foreground"
        />
        <button
          disabled={!valid}
          className="rounded p-2 border-2 border-foreground hover:bg-foreground/10 transition duration-300 ease-in-out disabled:opacity-50"
          onClick={install}
        >
          Install
        </button>
        <ErrorMessage message={errorMessage} />
      </div>
    </Modal>
  )
}
