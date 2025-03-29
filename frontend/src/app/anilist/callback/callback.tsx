'use client'

import { FullscreenCenter } from '@/components/FullscreenCenter'
import { useEffect, useState } from 'react'
import { AnilistLogin } from '../../../../wailsjs/go/backend/App'

export default function Callback () {
  const [accessToken, setAccessToken] = useState<string | null>(null)
  const [message, setMessage] = useState<string>('Logging in...')

  useEffect(() => {
    const token = window.location.hash.replace('#access_token=', '').replace(/&.*/, '')
    setAccessToken(token)
  }, [])

  useEffect(() => {
    if (accessToken) {
        AnilistLogin(accessToken)
            .then(() => { window.location.href = '/manga' })
            .catch((error) => {
              console.error(error)
              setMessage('Could not login to AniList')
            })
    }
  }, [accessToken])

  return (
    <FullscreenCenter>
      <p className="max-w-[50%] break-words">{message}</p>
    </FullscreenCenter>
  )
}
