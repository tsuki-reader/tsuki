import { Modal } from '../Modal'
import Image from 'next/image'
import { ProviderRow } from './ProviderRow'
import { useState } from 'react'
import { ErrorMessage } from '../ErrorMessage'
import { SuccessMessage } from '../SuccessMessage'
import { extensions } from '../../../wailsjs/go/models'
import { RepositoriesDestroy, RepositoriesUpdate } from '../../../wailsjs/go/backend/App'

interface Props {
    repository: extensions.Repository
    opened: boolean
    onClose: React.ReactEventHandler<HTMLDialogElement>
    onRepoUninstall: (repos: extensions.Repository[]) => void
    onRepoUpdate: (repo: extensions.Repository, idChanged: boolean, oldRepoId: string) => void
}

// TODO: Send a request to retrieve the providers
export function RepositoryModal ({ repository, opened, onClose, onRepoUninstall, onRepoUpdate }: Props) {
  const [errorMessage, setErrorMessage] = useState<string | null>(null)
  const [successMessage, setSuccessMessage] = useState<string | null>(null)
  const [mangaProviders, setMangaProviders] = useState<extensions.Provider[]>(repository.manga_providers)
  const [comicProviders, setComicProviders] = useState<extensions.Provider[]>(repository.comic_providers)

  const uninstall = () => {
    RepositoriesDestroy(repository.id)
        .then((repos) => onRepoUninstall(repos))
        .catch((error) => setErrorMessage(`There was an error uninstalling the repository: ${error}`))
  }

  const update = () => {
    setErrorMessage(null)
    setSuccessMessage(null)
    RepositoriesUpdate(repository.id)
        .then((repo: extensions.Repository) => {
            setErrorMessage(null)
            setSuccessMessage('Repository updated successfully')
            onRepoUpdate(repo, repo.id !== repository.id, repository.id)
        })
        .catch((error) => setErrorMessage(`There was an error updating the repository: ${error}`))
  }

  const closeModal = (event: React.SyntheticEvent<HTMLDialogElement, Event>) => {
    setErrorMessage(null)
    setSuccessMessage(null)
    onClose(event)
  }

  const onProviderInstalled = (providers: extensions.Provider[], type: string) => {
    if (type === 'comics') {
      setComicProviders(providers)
      repository.comic_providers = providers
    } else {
      setMangaProviders(providers)
      repository.manga_providers = providers
    }
    onRepoUpdate(repository, false, repository.id)
  }

  const providerRows = (providers: extensions.Provider[], providerType: string) => {
    if (providers.length === 0) {
      return (
        <div className="w-full text-center">
            No providers available
        </div>
      )
    }

    return providers.map((provider, index) => <ProviderRow key={index} provider={provider} repository={repository} providerType={providerType} onProviderInstalled={onProviderInstalled} />)
  }

  return (
        <Modal opened={opened} onClose={closeModal}>
            <div className="flex justify-center mb-8">
                <Image
                    className="h-16 w-auto"
                    src={repository.logo}
                    alt={`${repository.name} Logo`}
                    width={0}
                    height={0}
                />
            </div>
            <div className="flex flex-col w-full items-center mb-8">
                <h1 className="text-3xl font-bold">{repository.name}</h1>
                <small>{repository.id}</small>
            </div>
            <div className="flex flex-col gap-4 mb-8">
                <div className="flex flex-row w-full justify-center gap-4">
                    <button className="whitespace-nowrap h-fit rounded-full py-2 px-4 border-2 border-foreground hover:bg-foreground/10 transition duration-300 ease-in-out"
                        onClick={update}
                    >
                        Update
                    </button>
                    <button className="whitespace-nowrap h-fit rounded-full py-2 px-4 border-2 border-red-500 hover:bg-red-500/10 text-red-500 transition duration-300 ease-in-out"
                        onClick={uninstall}
                    >
                        Uninstall
                    </button>
                </div>
                <ErrorMessage message={errorMessage} />
                <SuccessMessage message={successMessage} />
            </div>
            <div className="flex flex-col w-full gap-8">
                <details open>
                    <summary className="font-bold text-xl cursor-pointer">Manga providers</summary>
                    <div className="mt-4 flex flex-col gap-2">
                        {providerRows(mangaProviders, 'manga')}
                    </div>
                </details>

                <details>
                    <summary className="font-bold text-xl cursor-pointer">Comic providers</summary>
                    <div className="mt-4 flex flex-col gap-2">
                        {providerRows(comicProviders, 'comics')}
                    </div>
                </details>
            </div>
        </Modal>
  )
}
