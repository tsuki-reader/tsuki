import { Account } from '@/types/models'
import { createContext } from 'react'

export const AccountContext = createContext<Account | null>(null)

export const AccountProvider = AccountContext.Provider
