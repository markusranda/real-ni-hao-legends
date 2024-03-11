
export interface NHScore {
	userId: string
	score: number
	townName: string
}

export interface NHSamba {
	scores: NHScore[] 
}
