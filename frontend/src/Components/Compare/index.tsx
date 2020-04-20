import React, { useState, useMemo, useEffect } from 'react';
import { useHistory } from 'react-router-dom';
import { ButtonBase } from '@aragon/ui';
import { Card, CardHeader, CardBody, Button } from 'shards-react';
import { getList, vote } from '../../api';
import CroppedImg from '../common/CroppedImg';
import CenterMessage from '../common/CenterMessage';

import type { food } from '../../types'

const MIN_PLAYED_RESULT = 5;
const MAX_REPEAT = 5;

function removeFromList(list: number[], items:number[]): number[] {
  return list.filter((element) => !items.includes(element));
}

function Compare() {
  const history = useHistory();

  // Can go to result after 10 plays
  const [countPlayed, setCountPlayed] = useState(0);
  const [isLoading, setIsLoading] = useState<boolean>(true)
  
  // const [winnderHistory, setWinnerHistory] = useState<number[]>([])
  const [winnerHistory, setWinnerHistory] = useState<number[]>([])

  const [done, setDone] = useState<boolean>(false);
  const [foodList, setFoodList] = useState<food[]>([]);
  const [indexList, setIndexList] = useState<number[]>([]);
  // index of the indexList
  const [optionAIdx, setOptionAIdx] = useState<number>(0);
  const [optionBIdx, setOptionBIdx] = useState<number>(0);

  // init states
  useMemo(async () => {
    const _foodList = await getList();
    const initIndexs = Array.from(Array(_foodList.length).keys()).sort(() => 0.5 - Math.random());
    const [a, b] = initIndexs.slice(0,2)
    
    setOptionAIdx(a);
    setOptionBIdx(b);
    setIndexList(removeFromList(initIndexs, [a, b]));
    setFoodList(_foodList);
    setIsLoading(false)
  }, []);

  useEffect(() => {
    if (!isLoading && indexList.length === 0) {
      setDone(true);
      return;
    }
  }, [indexList, isLoading])

  const updateBothIdxs = () => {
    if (indexList.length < 2) {
      setDone(true)
      return
    }
    const [newAIdx, newBIdx] = indexList.slice(0,2)
    setIndexList(indexList.slice(2))
    setOptionAIdx(newAIdx)
    setOptionBIdx(newBIdx)
  }

  const onClickOptionA = () => {
    
    setCountPlayed(countPlayed + 1);
    vote(foodList[optionAIdx].id, foodList[optionBIdx].id);
    setWinnerHistory(winnerHistory.concat(optionAIdx))

    if(winnerHistory.filter(idx => idx === optionAIdx).length > MAX_REPEAT) {
      updateBothIdxs()
    } else {
      const newOptionB = indexList[0];
      setIndexList(indexList.slice(1));
      setOptionBIdx(newOptionB);
    }
  };

  const onClickOptionB = () => {
    setCountPlayed(countPlayed + 1);
    vote(foodList[optionBIdx].id, foodList[optionAIdx].id);

    setWinnerHistory(winnerHistory.concat(optionBIdx))
    if(winnerHistory.filter(idx => idx === optionBIdx).length > MAX_REPEAT) {
      updateBothIdxs()
    } else {
      const newOptionA = indexList[0];
      setIndexList(indexList.slice(1));
      setOptionAIdx(newOptionA);
    }    
  };

  return <div>
      <div style={{ paddingTop: '6%', paddingBottom: '4%', textAlign: 'center', fontSize: 36 }}>
        哪個好吃
      </div>
      { isLoading ? (
        <CenterMessage text='載入中...' />
      ) : done ? (
        <CenterMessage text='沒東西啦！' />
      ) : (
        <div style={{ textAlign: 'center' }}>
          <div style={{ padding: '3%', display: 'inline-block' }}>
            <Card style={{ maxWidth: '650' }}>
              <ButtonBase disabled={done} onClick={onClickOptionA}>
                <CardHeader></CardHeader>
                <CroppedImg
                  width={350}
                  height={300}
                  url={
                    foodList.length === 0
                      ? 'https://place-hold.it/300x200'
                      : foodList[optionAIdx].image
                  }
                />
                <CardBody>
                  <div style={{ fontSize: 25 }}>
                    {foodList.length === 0 ? '' : foodList[optionAIdx].name}
                  </div>
                </CardBody>
              </ButtonBase>
            </Card>
          </div>
          <div style={{ padding: '3%', display: 'inline-block' }}>
            <Card style={{ maxWidth: '500' }}>
              <ButtonBase disabled={done} onClick={onClickOptionB}>
                <CardHeader></CardHeader>
                <CroppedImg
                  width={350}
                  height={300}
                  url={
                    foodList.length === 0
                      ? 'https://place-hold.it/300x200'
                      : foodList[optionBIdx].image
                  }
                />
                <CardBody>
                  <div style={{ fontSize: 25 }}>
                    {foodList.length === 0 ? '' : foodList[optionBIdx].name}
                  </div>
                </CardBody>
              </ButtonBase>
            </Card>
          </div>
        </div>
      )}

      <div style={{ textAlign: 'center' }}>
        <div style={{ padding: '3%', display: 'inline-block' }}>
          {countPlayed >= MIN_PLAYED_RESULT ? (
            <Button display='true' onClick={() => history.push('/ranking')}>
              {' '}
              看結果{' '}
            </Button>
          ) : (
            <></>
          )}
        </div>
      </div>
    </div>
  ;
}

export default Compare;
