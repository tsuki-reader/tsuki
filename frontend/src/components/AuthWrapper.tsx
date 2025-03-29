'use client'

import { useEffect, useState } from 'react'
import { LoadingScreen } from './LoadingScreen'
import { Login } from './auth/Login'
import { Account } from '@/types/models'
import { AccountProvider } from '@/contexts/account'

interface Props {
    children: React.ReactNode
}

export function AuthWrapper ({ children }: Props) {
  const [loading, setLoading] = useState(true)
  const [authenticated, setAuthenticated] = useState(false)
  const [account, setAccount] = useState<Account | null>(null)

  useEffect(() => {
    // Get the user from localStorage
    const user = localStorage.getItem('tsuki_account')
    if (user && user !== '') {
      setAccount(JSON.parse(user))
      setAuthenticated(true)
    } else {
      setAuthenticated(false)
    }
    setLoading(false)
  }, [])

  if (loading) {
    return <LoadingScreen />
  }

  if (authenticated && account) {
    return (
        <AccountProvider value={account}>
            {children}
        </AccountProvider>
    )
  } else {
    return <Login />
  }
}
