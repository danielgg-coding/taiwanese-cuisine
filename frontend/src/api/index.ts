import type { food } from '../types'


export const vote = (winner:number, loser:number) : void => {
  fetch(`/api/votef?loser=${loser}&winner=${winner}`)

};

export const getList = async () : Promise<food[]> => {
  const response = await fetch(`/api/cuisinef/`);
  const list = await response.json()
  return list
}
