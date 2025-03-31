export function LoadingState () {
  return (
    <div className="rounded h-full flex flex-col gap-4 justify-center items-center">
      <div className="z-10 w-[25px] h-[25px] loader border-[2px] border-t-[5px] border-foreground rounded-full animate-spin"></div>
    </div>
  )
}
