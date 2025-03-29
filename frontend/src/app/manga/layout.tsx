'use client'

import AnilistStatusWrapper from '@/components/anilist/AnilistStatusWrapper'

interface Props {
  children: React.ReactNode
}

export default function MangaLayout ({ children }: Props) {
  return (
    <AnilistStatusWrapper>
      {children}
    </AnilistStatusWrapper>
  )
}
