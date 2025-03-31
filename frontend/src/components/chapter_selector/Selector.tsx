'use client'

import { useEffect, useState } from 'react'
import { LoadingState } from './LoadingState'
import { ErrorState } from './ErrorState'
import { NoProviderState } from './NoProviderState'
import { NoProviders } from './NoProviders'
import Select from 'react-dropdown-select'
import styled from '@emotion/styled'
import { ChapterList } from './ChapterList'
import { AssignMapping, ProvidersIndex } from '../../../wailsjs/go/backend/App'
import { backend, models, types } from '../../../wailsjs/go/models'

interface Props {
  mediaList: types.ALMediaList
  initialChapters: models.Chapter[]
}

interface AssignResponse {
    mediaList: types.ALMediaList,
    chapters: models.Chapter[]
}

// TODO: Support both manga and comics
export function Selector ({ mediaList, initialChapters }: Props) {
  const [providers, setProviders] = useState<models.InstalledProvider[]>([])
  const [selectedProvider, setSelectedProvider] = useState<models.InstalledProvider | null>(mediaList.mapping?.installedProvider ?? null)
  const [errorMessage, setErrorMessage] = useState<string | null>(null)
  const [loading, setLoading] = useState(true)
  const [chapters, setChapters] = useState<models.Chapter[]>(initialChapters)

  useEffect(() => {
    const handleStatus = (providers: models.InstalledProvider[]) => {
      setProviders(providers)
      console.log(providers)
      setLoading(false)
    }

    const handleError = (e: {error: string}) => {
      setErrorMessage(e.error)
      setLoading(false)
    }

    // Retrieve list of providers
    ProvidersIndex("", "manga")
      .then(handleStatus)
      .catch(handleError)
  }, [])

  const currentState = () => {
    if (loading) {
      return <LoadingState />
    }

    if (errorMessage) {
      return <ErrorState message={errorMessage} />
    }

    if (providers.length === 0) {
      return <NoProviders />
    }

    if (selectedProvider === null) {
      return <NoProviderState />
    }

    return <ChapterList chapters={chapters} mediaList={mediaList} />
  }

  const handleAssign = (response: backend.MappingAssignResponse) => {
    setChapters(response.chapters)
  }

  const onProviderSelected = (newProvider: string | object | null) => {
    const p = newProvider !== null ? newProvider as models.InstalledProvider : null
    if (p === null) { return }
    setSelectedProvider(p)
    AssignMapping(mediaList.media.id.toString(), p.ID)
      .then(handleAssign)
      .catch((e: string) => setErrorMessage(`There was an error: ${e}`))
  }

  return (
    <div className="flex flex-col gap-4">
      <div className="w-full max-w-80">
        <StyledSelect
          options={providers}
          values={selectedProvider !== null ? [selectedProvider] : []}
          onChange={(value) => onProviderSelected(value.at(0) ?? null)}
          labelField="name"
          valueField="ID"
          searchBy="name"
          dropdownHandle={false}
          placeholder="Select a provider"
          color="#111427"
          dropdownHeight="150px"
        />
      </div>
      <div className="rounded min-h-96 bg-foreground/10 border-2 border-foreground">
        {currentState()}
      </div>
    </div>
  )
}

const StyledSelect = styled(Select)`
  border: 2px solid rgb(var(--foreground)) !important;
  border-radius: 0.25rem;
  padding: 5px 10px;
  background-color: rgba(var(--foreground-with-commas), 0.1);
  cursor: text;

  .react-dropdown-select-dropdown {
    border: 2px solid rgb(var(--foreground)) !important;
    border-radius: 0.25rem;
    background-color: rgb(var(--foreground));
    color: rgb(var(--background));
    left: -2px;
    top: 34px;
  }

  .react-dropdown-select-item-selected {
    background-color: rgb(var(--background)) !important;
    color: rgb(var(--foreground)) !important;
    font-weight: bold;
    border-bottom: none !important;
  }

  .react-dropdown-select-item {
    border-bottom: none !important;
  }

  .react-dropdown-select-input::placeholder {
    color: rgb(var(--foreground));
  }
`
