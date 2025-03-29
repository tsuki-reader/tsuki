'use client'

import Link from 'next/link'
import { types } from '../../../wailsjs/go/models'

interface Props {
  entry: types.ALMediaList
}

export function MangaPoster ({ entry }: Props) {
  const entryPercentage = () => {
    if (entry.media.chapters === 0) { return '0%' }
    return `${(entry.progress / entry.media.chapters) * 100}%`
  }

  const amountRemaining = () => {
    if (entry.media.chapters === 0) { return 0 }
    return entry.media.chapters - entry.progress
  }

  return (
    <Link href={`/manga/entry?id=${entry.media.id}`}>
      <div className="flex flex-col gap-3 cursor-pointer h-full group">
        <div className="p-2 group-hover:rotate-2 relative aspect-[6/8] transition-all duration-100 ease-in-out">
          <div className="relative col-span-1 w-full h-full bg-cover bg-center rounded overflow-hidden border-2 border-foreground" style={{ backgroundImage: `url(${entry.media.coverImage.large})` }}>
            <div className="absolute bottom-0 left-0 w-full h-2 bg-foreground" style={{ width: entryPercentage() }}></div>
            {amountRemaining() > 0 &&
              <div className="absolute text-sm flex justify-center items-center rounded-full bg-background text-foreground border-2 border-foreground p-2 m-2 top-0 right-0 h-8 min-w-8">
                {amountRemaining()}
              </div>
            }
          </div>
        </div>
        <div className="pt-0 p-2">
          <p className="text-md font-bold">{entry.media.title.english !== '' ? entry.media.title.english : entry.media.title.romaji}</p>
          <p className="text-sm font-thin">{entry.media.startDate.year}</p>
        </div>
      </div>
    </Link>
  )
}
