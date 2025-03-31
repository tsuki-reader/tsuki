'use client'

import { useSearchParams } from 'next/navigation'
import { FullscreenCenter } from '../FullscreenCenter'
import { useEffect, useState } from 'react'
import { LoadingScreen } from '../LoadingScreen'
import { ErrorMessage } from '../ErrorMessage'
import { MangaHeader } from './MangaHeader'
import { MangaShow } from '../../../wailsjs/go/backend/App'
import { backend, models, types } from '../../../wailsjs/go/models'
import { Selector } from '../chapter_selector/Selector'

export function MangaPage () {
  const [mediaList, setMediaList] = useState<types.ALMediaList | undefined>(undefined)
  const [chapters, setChapters] = useState<models.Chapter[]>([])
  const [errorMessage, setErrorMessage] = useState<string | null>(null)

  const searchParams = useSearchParams()
  const id = searchParams.get('id')

    useEffect(() => {
        const handleStatus = (response: backend.MangaShowResponse) => {
            setMediaList(response.media_list)
            setChapters(response.chapters)
        }

        const handleError = (e: string) => {
            setErrorMessage(e)
        }

        if (id && id.trim() !== '') {
            MangaShow(parseInt(id)).then(handleStatus).catch(handleError)
        } else {
            setErrorMessage('Manga entry not found.')
        }
    }, [id])

  if (mediaList === undefined && errorMessage === null) {
    return <LoadingScreen />
  }

  if (errorMessage) {
    return (
      <FullscreenCenter>
        <ErrorMessage message={errorMessage} />
      </FullscreenCenter>
    )
  }

  return (
    <>
      <div className="my-[150px] mx-12">
        <MangaHeader mediaList={mediaList!} />
        <Selector mediaList={mediaList!} initialChapters={chapters} />
      </div>
      {mediaList!.media.bannerImage !== '' &&
        <div className="bg-overlay fixed top-0 left-0 h-screen w-full bg-cover bg-no-repeat bg-center -z-[1]" style={{ backgroundImage: `url(${mediaList!.media.bannerImage})` }}></div>
      }
    </>
  )
}
