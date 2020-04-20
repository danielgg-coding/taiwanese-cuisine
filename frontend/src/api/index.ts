import type { food } from '../types'


export const vote = (winner:number, loser:number) : void => {
  fetch(`/votef?loser=${loser}&winner=${winner}`)

};

export const getList = async () : Promise<food[]> => {
  const response = await fetch(`/cuisinef/`);
  return response.json()
}
