'use client'

import { sanitizeText } from '@/helpers/text'
import { MediaList } from '@/types/anilist'
import { faCalendar } from '@fortawesome/free-solid-svg-icons'
import { FontAwesomeIcon } from '@fortawesome/react-fontawesome'
import Image from 'next/image'

interface Props {
  mediaList: MediaList
}

export function MangaHeader ({ mediaList }: Props) {
  const mainTitle = () => {
    return mediaList.media.title.english === '' ? mediaList.media.title.romaji : mediaList.media.title.english
  }

  const chapterCount = () => {
    if (mediaList.media.chapters === 0) { return mediaList.progress }
    return mediaList.media.chapters
  }

  return (
    <div className="flex flex-col gap-16 mb-16">
      <div className="flex gap-16">
        <Image width={250}
          height={0}
          src={mediaList.media.coverImage.large}
          alt={`${mediaList.media.title.romaji} Poster`}
          className="min-w-[250px] h-fit rounded border-2 border-foreground"
        />
        <div className="flex flex-col gap-4">
          <h1 className="font-bold text-4xl">
            {mainTitle()}
          </h1>
          {mainTitle() !== mediaList.media.title.romaji &&
            <p className="text-xl">{mediaList.media.title.romaji}</p>
          }
          <div className="flex flex-wrap gap-4">
            <div className="flex gap-2 items-center px-4 py-2 rounded-full bg-foreground/10 border-2 border-foreground w-fit">
              <FontAwesomeIcon icon={faCalendar} className="text-xs" />
              {mediaList.media.startDate.year}
            </div>
            <div className="flex gap-2 items-center px-4 py-2 rounded-full bg-foreground/10 border-2 border-foreground w-fit">
              {mediaList.media.status.toLowerCase()}
            </div>
            <div className="flex gap-2 items-center px-4 py-2 rounded-full bg-foreground/10 border-2 border-foreground w-fit">{mediaList.progress} / {chapterCount()}</div>
          </div>
          <div className="rounded bg-foreground/10 p-4 max-h-40 overflow-y-auto" dangerouslySetInnerHTML={{ __html: sanitizeText(mediaList.media.description) }}></div>
        </div>
      </div>
      <div className="flex flex-wrap gap-4">
        {mediaList.media.genres.map((genre, key) => <div key={key} className="px-2 py-1 rounded-full bg-foreground/10 border-2 border-foreground text-sm">{genre}</div>)}
      </div>
    </div>
  )
}
