import { faPuzzlePiece } from '@fortawesome/free-solid-svg-icons/faPuzzlePiece'
import { FontAwesomeIcon } from '@fortawesome/react-fontawesome'

export function NoProviderState() {
  return (
    <div className="rounded h-96 flex flex-col gap-4 justify-center items-center">
      <FontAwesomeIcon icon={faPuzzlePiece} className="text-4xl" />
      <p className="text-xl">No provider selected</p>
    </div>
  )
}
