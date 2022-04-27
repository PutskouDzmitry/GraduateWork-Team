import routerActionTypes from "../constants/routerActionTypes";

let initialState = {
  routersList: [],
  currentRouter: {
    id: "",
    settings: {
      val1: "",
      val2: 0,
    },
  },
};

const modalActionsReducer = (state = initialState, action) => {
  let newState = {};

  switch (action.type) {
    case routerActionTypes.ADD_ROUTER:
      newState = {
        ...state,
        routersList: [
          ...state.routersList,
          {
            id: action.id,
            coords: action.coords,
            settings: action.settings,
          },
        ],
      };
      return newState;

    case routerActionTypes.UPDATE_ROUTER:
      newState = {
        ...state,
        routersList: state.routersList.map((router) => {
          return router.id == action.id
            ? {
                ...router,
                settings: action.settings,
              }
            : router;
        }),
      };
      return newState;

    case routerActionTypes.REMOVE_ROUTER:
      newState = {
        ...state,
        routersList: state.routersList.filter((router) => {
          return router.id != action.id;
        }),
      };
      console.log(newState);
      return newState;

    case routerActionTypes.SET_CURRENT_ROUTER_ID:
      newState = {
        ...state,
        currentRouter: {
          ...state.currentRouter,
          id: action.id,
        },
      };
      return newState;

    case routerActionTypes.SET_CURRENT_ROUTER_SETTINGS:
      newState = {
        ...state,
        currentRouter: {
          ...state.currentRouter,
          settings: action.settings,
        },
      };
      return newState;

    default:
      return state;
  }
};

export default modalActionsReducer;
