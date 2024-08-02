import { createContext } from 'react';

function InitMenuContext<T>(initialData: T) {
  const menuContext = createContext<T>(initialData);

  return menuContext;
}
