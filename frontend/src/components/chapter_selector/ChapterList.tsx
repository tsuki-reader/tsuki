'use client'

import { capitalizeText } from '@/helpers/text'
import { MediaList } from '@/types/anilist'
import { useEffect, useState } from 'react'
import Reader from '../Reader/Reader'
import { models } from '../../../wailsjs/go/models'

interface Props {
  mediaList: MediaList
  chapters: models.Chapter[]
}

export function ChapterList ({ mediaList, chapters }: Props) {
  const [page, setPage] = useState(1)
  const [currentChapters, setCurrentChapters] = useState<models.Chapter[]>([])
  const [maxPages, setMaxPages] = useState(1)
  const [currentChapter, setCurrentChapter] = useState<models.Chapter | undefined>(undefined)

  useEffect(() => {
    const max = Math.ceil(chapters.length / 10)
    setMaxPages(max)
  }, [chapters])

  useEffect(() => {
    const start = (page - 1) * 10
    const end = start + 10
    const newChapters = chapters.slice(start, end)
    setCurrentChapters(newChapters)
  }, [page, chapters])

  const getPaginationRange = () => {
    if (maxPages <= 4) return Array.from({ length: maxPages }, (_, i) => i + 1)

    if (page <= 2) return [1, 2, 3, 4]
    if (page >= maxPages - 2) return [maxPages - 3, maxPages - 2, maxPages - 1, maxPages]

    return [page - 2, page - 1, page, page + 1]
  }

  return (
    <>
      <table className="min-w-full divide-y divide-navy bg-transparent text-sm">
        <thead className="text-left">
          <tr>
            <th className="px-4 py-4 font-bold whitespace-nowrap">Name</th>
            <th className="px-4 py-4 font-bold whitespace-nowrap w-0">Absolute number</th>
          </tr>
        </thead>

        <tbody className="divide-y divide-navy">
          {currentChapters.map((chapter, index) => (
              <tr className="cursor-pointer hover:bg-foreground/30" key={index} onClick={() => { setCurrentChapter(chapter) }}>
              <td className="px-4 py-4 font-medium whitespace-nowrap">{capitalizeText(chapter.title)}</td>
              <td className="px-4 py-4 whitespace-nowrap">{chapter.absolute_number}</td>
            </tr>
          ))}
        </tbody>
      </table>

      <div className="rounded-b-lg border-t border-foreground px-4 py-2">
        <ol className="flex justify-end gap-1 text-xs font-medium">
          <li>
            <button
              className="cursor-pointer inline-flex size-8 items-center justify-center rounded-sm border border-foreground"
              onClick={() => { setPage(1) }}>
              <span className="sr-only">First Page</span>
              <svg
                xmlns="http://www.w3.org/2000/svg"
                className="size-3"
                viewBox="0 0 20 20"
                fill="currentColor">
                <path
                  fillRule="evenodd"
                  d="M12.707 5.293a1 1 0 010 1.414L9.414 10l3.293 3.293a1 1 0 01-1.414 1.414l-4-4a1 1 0 010-1.414l4-4a1 1 0 011.414 0z"
                  clipRule="evenodd" />
              </svg>
            </button>
          </li>

          {getPaginationRange().map((pageNum) => (
            <li key={pageNum}>
              <button
                className={`cursor-pointer block size-8 rounded-sm border border-foreground text-center leading-8 ${page === pageNum ? 'bg-foreground text-background' : ''}`}
                onClick={() => setPage(pageNum)}>
                  {pageNum}
              </button>
            </li>
          ))}

          <li>
            <button
              className="cursor-pointer inline-flex size-8 items-center justify-center rounded-sm border border-foreground"
              onClick={() => { setPage(maxPages) }}>
              <span className="sr-only">Last Page</span>
              <svg
                xmlns="http://www.w3.org/2000/svg"
                className="size-3"
                viewBox="0 0 20 20"
                fill="currentColor">
                <path
                  fillRule="evenodd"
                  d="M7.293 14.707a1 1 0 010-1.414L10.586 10 7.293 6.707a1 1 0 011.414-1.414l4 4a1 1 0 010 1.414l-4 4a1 1 0 01-1.414 0z"
                  clipRule="evenodd" />
              </svg>
            </button>
          </li>
        </ol>
      </div>

      <Reader mediaList={mediaList} currentChapter={currentChapter} setCurrentChapter={setCurrentChapter} />
    </>
  )
}
