import stepActionTypes from "../constants/stepActionTypes";

let initialState = {
  stepsList: [],
};

const modalActionsReducer = (state = initialState, action) => {
  let newState = {};

  switch (action.type) {
    case stepActionTypes.ADD_STEP:
      newState = {
        ...state,
        stepsList: [
          ...state.stepsList,
          {
            id: action.id,
            coords: action.coords,
          },
        ],
      };
      return newState;

    case stepActionTypes.REMOVE_STEP:
      newState = {
        ...state,
        stepsList: state.stepsList.filter((step) => {
          return step.id != action.id;
        }),
      };
      return newState;

    case stepActionTypes.REMOVE_ALL_STEPS:
      newState = {
        ...state,
        stepsList: [],
      };
      return newState;

    default:
      return state;
  }
};

export default modalActionsReducer;
