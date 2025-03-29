'use client'

import React from 'react'

interface Props {
    children: React.ReactNode
}

export function FullscreenCenter ({ children }: Props) {
  return (
        <div className="flex flex-col justify-center items-center w-full min-h-screen">
            {children}
        </div>
  )
}
