
export const vote = async (winner, loser) => {
  fetch(`/api/votef?loser=${loser}&winner=${winner}`).then(res => {
    console.log(res)
  }).catch(error => {
    console.error(error)
  })

};

/**
 * @return {Promise<{Name:string, Image:string, Played: number, Score: number}[]>}
 */
export const getList = async () => {
  const response = await fetch(`/api/cuisinef/`);
  const list = await response.json()
  return list
}
