import { NHQuizQuestion } from '@/models/nh-quiz'

export const question_1: NHQuizQuestion = {
  id: 'question_1',
  question: 'Who is the strongest snake?',
  answers: [
    {
      answer: 'Honey badger',
      question: 'question_2'
    },
    {
      answer: 'King cobra',
      question: 'question_1'
    },
    {
      answer: 'Anaconda',
      question: 'question_1'
    },
    {
      answer: 'Onyx',
      question: 'question_1'
    }
  ]
}

const question_2: NHQuizQuestion = {
  id: 'question_2',
  question: 'Do you have turkey skin?',
  answers: [
    {
      answer: 'No',
      question: 'question_3'
    },
    {
      answer: 'Yes',
      question: 'question_1'
    }
  ]
}

const question_3: NHQuizQuestion = {
  id: 'question_3',
  question: 'What is the best pokemon?',
  answers: [
    {
      answer: 'Pikachu',
      question: 'question_1'
    },
    {
      answer: 'Charizard',
      question: 'question_4'
    },
    {
      answer: 'Ferrothorn',
      question: 'question_4'
    },
    {
      answer: 'Cocoon',
      question: 'question_4'
    }
  ]
}

const question_4: NHQuizQuestion = {
  id: 'question_4',
  question: 'What is astrology?',
  answers: [
    {
      answer: 'A fraud',
      question: 'question_5'
    },
    {
      answer: 'The truth',
      question: 'question_1'
    },
    {
      answer: 'No',
      question: 'question_5'
    },
    {
      answer: 'Yes',
      question: 'question_1'
    }
  ]
}

const question_5: NHQuizQuestion = {
  id: 'question_5',
  question: 'What is the next question?',
  answers: [
    {
      answer: '6',
      question: 'question_6'
    },
    {
      answer: '5',
      question: 'question_1'
    },
    {
      answer: '4',
      question: 'question_1'
    },
    {
      answer: '7',
      question: 'question_1'
    }
  ]
}

const question_6: NHQuizQuestion = {
  id: 'question_6',
  question: 'Select right',
  answers: [
    {
      answer: 'right',
      question: 'question_7'
    },
    {
      answer: 'left',
      question: 'question_1'
    }
  ]
}

const question_7: NHQuizQuestion = {
  id: 'question_7',
  question: 'When was Erika Nissen born?',
  answers: [
    {
      answer: '13. jan. 1845',
      question: 'question_1'
    },
    {
      answer: '14. jan. 1845',
      question: 'question_1'
    },
    {
      answer: '15. jan. 1845',
      question: 'question_1'
    },
    {
      answer: '16. jan. 1845',
      question: 'question_8'
    },
    {
      answer: '17. jan. 1845',
      question: 'question_1'
    },
    {
      answer: '18. jan. 1845',
      question: 'question_1'
    }
  ]
}

const question_8: NHQuizQuestion = {
  id: 'question_8',
  question: 'What is the Randator method?',
  answers: [
    {
      answer: 'A festive dance',
      question: 'question_1'
    },
    {
      answer: 'A way to calculate the resistance in a electronics circuit',
      question: 'question_1'
    },
    {
      answer: 'Potato peeling method',
      question: 'question_1'
    },
    {
      answer: 'Technique used to balance chemical equations',
      question: 'question_9'
    }
  ]
}

const question_9: NHQuizQuestion = {
  id: 'question_9',
  question: 'The best number of explosions',
  answers: [
    {
      answer: '1',
      question: 'question_1'
    },
    {
      answer: '2',
      question: 'question_1'
    },
    {
      answer: '3',
      question: 'question_1'
    },
    {
      answer: '4',
      question: 'question_1'
    },
    {
      answer: '5',
      question: 'question_1'
    },
    {
      answer: '6',
      question: 'question_1'
    },
    {
      answer: '7',
      question: 'question_10'
    },
    {
      answer: '8',
      question: 'question_1'
    },
    {
      answer: '9',
      question: 'question_1'
    }
  ]
}

const question_10: NHQuizQuestion = {
  id: 'question_10',
  question: 'Best host of show',
  answers: [
    {
      answer: 'Fridtjof Wilborn',
      finished: true
    },
    {
      answer: 'Dan Børge Akerø',
      question: 'question_1'
    },
    {
      answer: 'Emil Andre Valde',
      question: 'question_1'
    },
    {
      answer: 'John Almås',
      question: 'question_1'
    },
    {
      answer: 'Fredrik Skavlan',
      question: 'question_1'
    },
    {
      answer: 'Elsa Kossheadset Furuseth',
      question: 'question_1'
    },
    {
      answer: 'Tore på sporet',
      question: 'question_1'
    }
  ]
}

export const QUESTIONS = {
  question_1,
  question_2,
  question_3,
  question_4,
  question_5,
  question_6,
  question_7,
  question_8,
  question_9,
  question_10
} as Record<string, NHQuizQuestion>
