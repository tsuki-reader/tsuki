'use client'

import { useEffect, useState } from 'react'
import { RepositoryButton } from './RepositoryButton'
import { InstallRepository } from './InstallRepository'
import { RepositoryModal } from './RepositoryModal'
import { RepositoriesIndex } from '../../../wailsjs/go/backend/App'
import { extensions } from '../../../wailsjs/go/models'

export function Repositories () {
  const [repositories, setRepositories] = useState<extensions.Repository[]>([])
  const [selectedRepo, setSelectedRepo] = useState<extensions.Repository | null>(null)
  const [installRepoOpened, setInstallRepoOpened] = useState<boolean>(false)

  useEffect(() => {
    RepositoriesIndex().then(handleResponse).catch((error) => console.log(error))
  }, [])

  const handleResponse = (repos: extensions.Repository[]) => {
    setRepositories(repos)
  }

  const renderRepositories = () => {
    return repositories.map((repository, index) => <RepositoryButton key={index} repository={repository} onClick={() => openRepoModal(repository)} />)
  }

  const openInstallRepoModal = () => {
    setInstallRepoOpened(true)
  }

  const closeInstallRepoModal = () => {
    setInstallRepoOpened(false)
  }

  const openRepoModal = (repo: extensions.Repository) => {
    setSelectedRepo(repo)
  }

  const closeRepoModal = () => {
    setSelectedRepo(null)
  }

  const onRepoInstalled = (repos: extensions.Repository[]) => {
    setInstallRepoOpened(false)
    setRepositories(repos)
  }

  const onRepoUninstalled = (repos: extensions.Repository[]) => {
    setSelectedRepo(null)
    setRepositories(repos)
  }

  const onRepoUpdated = (repo: extensions.Repository, idChanged: boolean, oldRepoId: string) => {
    let repos: extensions.Repository[] = []
    if (idChanged) {
      repos = repositories.filter((repository) => repository.id !== oldRepoId)
    } else {
      repos = repositories.filter((repository) => repository.id !== repo.id)
    }

    setRepositories([...repos, repo].sort((repoA, repoB) => repoA.id.localeCompare(repoB.id)))
    setSelectedRepo(repo)
  }

  return (
        <div className="my-[150px] mx-12">
            <div className="flex justify-between">
                <h1 className="text-3xl font-bold">Your repositories</h1>
                <button className="whitespace-nowrap h-fit rounded-full py-2 px-4 border-2 border-foreground hover:bg-foreground/10 transition duration-300 ease-in-out"
                    onClick={openInstallRepoModal}
                >
                    Add a repository
                </button>
            </div>
            <div className="py-8 flex flex-row gap-4 flex-wrap">
                {renderRepositories()}
            </div>
            <InstallRepository opened={installRepoOpened} onRepoInstalled={onRepoInstalled} onClose={closeInstallRepoModal} />
            {selectedRepo !== null &&
                <RepositoryModal repository={selectedRepo}
                    opened={selectedRepo !== null}
                    onClose={closeRepoModal}
                    onRepoUninstall={onRepoUninstalled}
                    onRepoUpdate={onRepoUpdated}
                />
            }
        </div>
  )
}
