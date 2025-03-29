'use client'

import { MangaPoster } from './MangaPoster'
import { types } from '../../../wailsjs/go/models'
import { useRef } from 'react'

interface Props {
    list: types.ALMediaListGroup
    open?: boolean
}

export function MediaListDisclosure ({ list, open }: Props) {
    const detailsRef = useRef<HTMLDetailsElement>(null)

    const handleClick = () => {
        if (detailsRef.current) {
            detailsRef.current.open = !detailsRef.current.open
        }
    }

    return (
        <div>
            <div className="cursor-pointer font-bold text-4xl w-fit rounded bg-foreground/10 p-4" onClick={handleClick}>
                {list.name}
            </div>
            <details ref={detailsRef} open={open}>
                <summary className='hidden'></summary>
                <div className="pt-8 gap-6 grid grid-cols-2 min-[768px]:grid-cols-3 min-[1080px]:grid-cols-4 min-[1320px]:grid-cols-5 min-[1750px]:grid-cols-6 min-[1850px]:grid-cols-7 min-[2000px]:grid-cols-8">
                    {list.entries.map((entry, key) => <MangaPoster key={key} entry={entry} />)}
                </div>
            </details>
        </div>
    )
}
