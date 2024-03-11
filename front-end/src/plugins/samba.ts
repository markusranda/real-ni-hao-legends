export const DANCE_MOVES_CIRCLE = {
    name: 'CIRCLE',
    moves: [
      [
        [-1, 0],
        [-1, 0]
      ],
      [
        [0, -1],
        [0, -1]
      ],
      [
        [1, 0],
        [1, 0]
      ],
      [
        [0, 1],
        [0, 1]
      ]
    ]
  }
  
  export const DANCE_MOVES_JACKS = {
    name: 'JACK',
    moves: [
      [
        [1, 0],
        [-1, 0]
      ],
      [
        [1, 0],
        [-1, 0]
      ],
      [
        [-1, 0],
        [1, 0]
      ],
      [
        [-1, 0],
        [1, 0]
      ]
    ]
  }
  
  export const DANCE_MOVES_TELEMARK = {
    name: 'TELEMARK',
    moves: [
      [
        [0, -1],
        [0, 1]
      ],
      [
        [0, -1],
        [0, 1]
      ],
  
      [
        [0, 1],
        [0, -1]
      ],
      [
        [0, 1],
        [0, -1]
      ]
    ]
  }
  
  export const DANCE_MOVES_TWIST = {
    name: 'TWIST',
    moves: [
      [
        [0, -1],
        [0, 1]
      ],
      [
        [0, 1],
        [0, -1]
      ]
    ]
  }
  
  export const DANCE_MOVES_FEETSPIN = {
    name: 'FEETSPIN',
    moves: [
      [
        [1, 0],
        [0, 1]
      ],
      [
        [0, 1],
        [-1, 0]
      ],
      [
        [-1, 0],
        [0, -1]
      ],
      [
        [0, -1],
        [1, 0]
      ]
    ]
  }
  
  export const DANCE_MOVES_SUMO_WRESTLER = {
    name: 'SUMO WRESTLER',
    moves: [
      [
        [2, 0],
        [0, 0]
      ],
      [
        [0, 0],
        [-2, 0]
      ],
      [
        [-2, 0],
        [2, 0]
      ]
    ]
  }
  
  export const DANCE_MOVES_TANGO = {
    name: 'TANGO',
    moves: [
      [
        [0, -1],
        [0, 1]
      ],
      [
        [0, 1],
        [0, -1]
      ],
      [
        [1, 0],
        [-1, 0]
      ],
      [
        [-1, 0],
        [1, 0]
      ]
    ]
  }
  
  export const DANCE_MOVES_MOONWALK = {
    name: 'MOONWALK',
    moves: [
      [
        [0, -3],
        [0, -3]
      ],
      [
        [0, 1],
        [0, 0]
      ],
      [
        [0, 0],
        [0, 2]
      ],
      [
        [0, 2],
        [0, 0]
      ],
      [
        [0, 0],
        [0, 2]
      ],
      [
        [0, 2],
        [0, 0]
      ],
      [
        [0, -2],
        [0, -1]
      ]
    ]
  }
  
  export const DANCE_MOVES_PIROUETTE = {
    name: 'PIRUETTE',
    moves: [
      [
        [2, -1],
        [-2, -1]
      ],
      [
        [0, 1],
        [0, -1]
      ],
      [
        [-1, 1],
        [1, -1]
      ],
      [
        [-1, 0],
        [1, 0]
      ],
      [
        [-1, 0],
        [1, 0]
      ],
      [
        [-1, -1],
        [1, 1]
      ],
      [
        [0, -1],
        [0, 1]
      ],
      [
        [2, 1],
        [-2, 1]
      ]
    ]
  }

export const DANCE_MOVES = [
    DANCE_MOVES_TELEMARK,
    DANCE_MOVES_CIRCLE,
    DANCE_MOVES_JACKS,
    DANCE_MOVES_TWIST,
    DANCE_MOVES_FEETSPIN,
    DANCE_MOVES_SUMO_WRESTLER,
    DANCE_MOVES_TANGO,
    DANCE_MOVES_MOONWALK,
    DANCE_MOVES_PIROUETTE
  ]