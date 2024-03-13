export interface NHCommand extends NHSendCommand{
  userId: string
  time: number
}

export interface NHSendCommand {
  type: string
  data?: any
}
