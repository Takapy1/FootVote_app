import React, { useEffect, useState } from 'react';
import axios from 'axios'

const GamesList = () => {
  const [games, setGames] = useState([])

  const getGames = () => {
    axios.get("/api/games")
    .then(resp => {
      setGames(resp.data) //"data": [Gameのリスト]で飛ばしてくる
    })
    .catch(e => {
      console.log(e)
    })
  }

  useEffect(() => {
    getGames()
  }, [])

  const GamesList = () => {
    games.map ((game) => {
      return (
        <tr>
          <td>{game.id}</td>
          <td>{game.my_team}</td>
          <td>{game.enemy_team}</td>
          <td>{game.date}</td>
        </tr>
      )
    })
  }


  return (
    <div>
      <table>
        <thead>
          <th>ID</th>
          <th>my_team</th>
          <th>enemy_team</th>
          <th>date</th>
        </thead>
        <tbody>
          <GamesList />
        </tbody>
      </table>
    </div>
  )
}

export default GamesList