'use client'

import { faDownload, faRotate, faTrashCan } from '@fortawesome/free-solid-svg-icons'
import { FontAwesomeIcon } from '@fortawesome/react-fontawesome'
import Image from 'next/image'
import { useState } from 'react'
import { extensions } from '../../../wailsjs/go/models'
import { ProvidersCreateOrUpdate, ProvidersDestroy } from '../../../wailsjs/go/backend/App'

interface Props {
    provider: extensions.Provider
    repository: extensions.Repository
    providerType: string
    onProviderInstalled: (providers: extensions.Provider[], type: string) => void
}

export function ProviderRow ({ provider, repository, providerType, onProviderInstalled }: Props) {
  const [loading, setLoading] = useState<boolean>(false)

  const handleResponse = (providers: extensions.Provider[]) => {
    onProviderInstalled(providers, providerType)
    setLoading(false)
  }

  const installProvider = () => {
    execBinding(ProvidersCreateOrUpdate)
  }

  const deleteProvider = () => {
    execBinding(ProvidersDestroy)
  }

  const execBinding = (binding: (repositoryId: string, providerId: string, providerType: string) => Promise<extensions.Provider[]>) => {
    setLoading(true)
    binding(repository.id, provider.id, providerType)
        .then(handleResponse)
        .catch((error) => {
          console.log(error)
          setLoading(false)
        })
  }

  return (
        <div className="w-full flex flex-row justify-between items-center p-2 border-2 border-foreground rounded bg-foreground/10">
            <div className="flex flex-row gap-2">
                <Image
                    className="h-6 w-auto"
                    src={provider.icon}
                    alt={`${provider.name} Logo`}
                    width={0}
                    height={0}
                />
                <p>{provider.name}</p>
            </div>
            {!provider.installed && !loading &&
                <FontAwesomeIcon onClick={installProvider} className="text-lg cursor-pointer text-green-500" icon={faDownload} />
            }
            {provider.installed && !loading &&
                (
                    <div className="flex flex-row gap-2">
                        <FontAwesomeIcon onClick={installProvider} className="text-lg cursor-pointer text-green-500" icon={faRotate} title="Update" />
                        <FontAwesomeIcon onClick={deleteProvider} className="text-lg cursor-pointer text-red-500" icon={faTrashCan} title="Uninstall" />
                    </div>
                )
            }
        </div>
  )
}
