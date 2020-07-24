import { GntInstance } from "../../api/models";
import { useState } from "react";

export interface FilterFields {
  name: boolean;
  primaryNode: boolean;
  secondaryNodes: boolean;
}

interface State {
  filter: string;
  filterFields: FilterFields;
}

const initialState: State = {
  filter: "",
  filterFields: {
    name: true,
    primaryNode: true,
    secondaryNodes: false,
  },
};

export const filterInstances = (
  instances: GntInstance[],
  filter: string,
  filterFields: FilterFields
): GntInstance[] => {
  if (filter === "") {
    return instances;
  }

  return instances.filter((instance) => {
    if (filterFields.name && instance.name.includes(filter)) {
      return true;
    }

    if (filterFields.primaryNode && instance.primaryNode.includes(filter)) {
      return true;
    }

    if (filterFields.secondaryNodes) {
      for (const node of instance.secondaryNodes) {
        if (node.includes(filter)) {
          return true;
        }
      }
    }

    return false;
  });
};

export const useFilter = (): [
  { filter: string; filterFields: FilterFields },
  {
    setFilter: (filter: string) => void;
    setFilterFields: (fields: FilterFields) => void;
    reset: () => void;
  }
] => {
  const [state, setState] = useState<State>(initialState);

  return [
    state,
    {
      setFilter: (filter) =>
        setState({
          filter,
          filterFields: state.filterFields,
        }),
      setFilterFields: (filterFields) =>
        setState({
          filter: state.filter,
          filterFields,
        }),
      reset: () => setState(initialState),
    },
  ];
};
