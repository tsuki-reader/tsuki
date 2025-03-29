import sanitizeHtml from 'sanitize-html'

export function sanitizeText (text: string): string {
  const defaultOptions: sanitizeHtml.IOptions = {
    allowedTags: ['br']
  }
  return sanitizeHtml(text, defaultOptions)
}

export function capitalizeText (text: string): string {
  const lower = text.toLowerCase()
  const words = lower.split(' ')
  const processedWords = words.map((word) => {
    return word.charAt(0).toUpperCase() + word.slice(1)
  })
  return processedWords.join(' ')
}
