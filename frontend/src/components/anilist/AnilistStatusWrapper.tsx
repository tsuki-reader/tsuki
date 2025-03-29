'use client'

import React, { useContext, useEffect, useState } from 'react'
import { AnilistLogin } from './AnilistLogin'
import { LoadingScreen } from '../LoadingScreen'
import { Viewer } from '@/types/anilist'
import { ViewerProvider } from '@/contexts/viewer'
import { AccountContext } from '@/contexts/account'
import { AnilistStatus } from '../../../wailsjs/go/backend/App'
import { backend } from '../../../wailsjs/go/models'

interface Props {
  children: React.ReactNode
}

export default function AnilistStatusWrapper ({ children }: Props) {
  const [loading, setLoading] = useState(true)
  const [authenticated, setAuthenticated] = useState(false)
  const [currentViewer, setCurrentViewer] = useState<Viewer | null>(null)
  const [clientId, setClientId] = useState<string>('')

  const account = useContext(AccountContext)

  const handleStatus = (status: backend.AnilistStatus) => {
    setClientId(status.client_id)
    if (status.authenticated && status.viewer !== null) {
      setAuthenticated(true)
      const viewer: Viewer = {
          name: status.viewer!.name,
          bannerImage: status.viewer!.bannerImage,
          avatar: status.viewer!.avatar,
      }
      setCurrentViewer(viewer)
    } else {
      setAuthenticated(false)
    }
    setLoading(false)
  }

  // eslint-disable-next-line @typescript-eslint/no-explicit-any
  const handleError = (error: any) => {
    console.error(error)
    setAuthenticated(false)
    setLoading(false)
  }

  useEffect(() => {
    AnilistStatus().then(handleStatus).catch(handleError)
  }, [account])

  if (loading) {
    return (
      <LoadingScreen />
    )
  }

  if (authenticated) {
    return (
      <ViewerProvider value={currentViewer}>
        {children}
      </ViewerProvider>
    )
  } else {
    return (
      <AnilistLogin clientId={clientId} />
    )
  }
}
