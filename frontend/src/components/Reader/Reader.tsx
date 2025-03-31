'use client'

import { MediaList } from '@/types/anilist'
import { faCircleXmark } from '@fortawesome/free-solid-svg-icons'
import { faGear } from '@fortawesome/free-solid-svg-icons/faGear'
import { FontAwesomeIcon } from '@fortawesome/react-fontawesome'
import { useEffect, useRef, useState } from 'react'
import { FullscreenCenter } from '../FullscreenCenter'
import { ErrorMessage } from '../ErrorMessage'
import { backend, models, providers } from '../../../wailsjs/go/models'
import { MangaChapterPages } from '../../../wailsjs/go/backend/App'

interface Props {
  mediaList: MediaList
  currentChapter: models.Chapter | undefined
  setCurrentChapter: (chapter: models.Chapter | undefined) => void
}

// TODO: Custom settings
export default function Reader ({ mediaList, currentChapter, setCurrentChapter }: Props) {
  const [open, setOpen] = useState<boolean | undefined>(undefined)
  const [currentPage, setCurrentPage] = useState<number>(1)
  const [errorMessage, setErrorMessage] = useState<string | null>(null)
  const [pages, setPages] = useState<providers.Page[]>([])
  // TODO: Fucking hack cos fuck this framework.
  const [key, setKey] = useState(0)
  const [immersive, setImmersive] = useState(false)

  const containerRef = useRef<HTMLDivElement>(null)

  const handleKeyDown = (event: KeyboardEvent) => {
    switch (event.key) {
      case 'Escape':
        closeReader()
        break
      case 'ArrowLeft':
        setCurrentPage(prev => (prev < pages.length ? prev + 1 : prev))
        break
      case 'ArrowRight':
        setCurrentPage(prev => (prev > 1 ? prev - 1 : prev))
        break
    }
  }

  useEffect(() => {
    setKey(prev => prev + 1)
  }, [currentPage])

  useEffect(() => {
    setCurrentPage(1)
    document.addEventListener('keydown', handleKeyDown)
    return () => document.removeEventListener('keydown', handleKeyDown)
  }, [pages])

  useEffect(() => {
    const handleStatus = (data: backend.MappingChapterPagesResponse) => {
      setPages(data.pages)
    }

    const handleError = (e: {error: string}) => {
      setErrorMessage(e.error)
    }

    if (currentChapter) {
      if (containerRef.current) containerRef.current.style.visibility = 'visible'
      setOpen(true)
      setCurrentPage(1)
      document.body.style.overflow = 'hidden'
      MangaChapterPages(mediaList.media.id, currentChapter.external_id)
        .then(handleStatus)
        .catch(handleError)
    } else {
      setOpen(false)
      document.body.style.overflow = 'auto'
    }
  }, [currentChapter])

  const closeReader = () => {
    setCurrentChapter(undefined)
    setPages([])
  }

  if (errorMessage) {
    return (
      <FullscreenCenter>
        <ErrorMessage message={errorMessage} />
      </FullscreenCenter>
    )
  }

  return (
    <div ref={containerRef} data-open={open} style={{ visibility: 'hidden' }} className='bg-black fixed left-0 inset-y-0 z-50 h-dvh w-full flex flex-col data-[open=true]:pointer-events-auto data-[open=false]:pointer-events-none data-[open=true]:animate-in data-[open=true]:slide-in-from-left data-[open=false]:animate-out data-[open=false]:slide-out-to-left transition fill-mode-forwards duration-300 ease-in-out'>
      <div className='absolute top-0 bottom-0 left-0 right-0 m-auto z-50 w-[500px] h-[500px] bg-transparent' onClick={() => setImmersive(prev => !prev)} />
      {!immersive && (
        <div className='w-full flex gap-2 bg-background'>
          <h2 className='m-2 font-bold'>{mediaList.media.title.english} - {currentChapter?.title}</h2>
        </div>
      )}

      <div key={key} className='w-full h-full flex-auto overflow-x-auto'>
        {pages.length > 0 &&
          <img src={pages[currentPage - 1].image_url} alt="Chapter Page" className='w-full object-contain' />
        }
      </div>

      {!immersive && (
        <div className='w-full flex gap-2 bg-background'>
          <FontAwesomeIcon icon={faCircleXmark} onClick={closeReader} className='text-2xl cursor-pointer m-4' />

          <div className='w-full m-4 ml-0 bg-foreground/25 h-2 my-auto rounded-full'>
            <div className='h-full bg-foreground rounded-full float-end' style={{ width: `${currentPage / Math.max(pages.length, 1) * 100}%` }}></div>
          </div>

          <FontAwesomeIcon icon={faGear} className='text-2xl cursor-pointer m-4' />
        </div>
      )}
    </div>
  )
}
