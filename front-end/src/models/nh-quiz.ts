export interface NHQuizAnswer {
  answer: string
  question?: string
  finished?: boolean
}

export interface NHQuizQuestion {
  id: string
  question: string
  answers: NHQuizAnswer[]
}
