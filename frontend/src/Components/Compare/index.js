import React, { useState, useMemo } from 'react';
import { useHistory } from 'react-router-dom';
import { ButtonBase } from '@aragon/ui';
import { Card, CardHeader, CardBody, Button } from 'shards-react';
import { getList, vote } from '../../api';
import CroppedImg from '../common/CroppedImg';
import CenterMessage from '../common/CenterMessage';

const MIN_PLAYED_RESULT = 8;

function chooseFromList(targetList) {
  const random = targetList[Math.floor(Math.random() * targetList.length)];
  return random;
}

function removeFromList(list, item) {
  return list.filter((element) => element !== item);
}

function initialAB(indexs) {
  const shuffled = indexs.sort(() => 0.5 - Math.random());
  let selected = shuffled.slice(0, 2);
  return selected;
}

function Compare() {
  const history = useHistory();

  // Can go to result after 10 plays
  const [countPlayed, setCountPlayed] = useState(0);

  const [done, setDone] = useState(false);
  const [foodList, setFoodList] = useState([]);
  const [indexList, updateIndexList] = useState([]);
  const [optionAIdx, setOptionAIdx] = useState(0);
  const [optionBIdx, setOptionBIdx] = useState(0);

  useMemo(async () => {
    const _foodList = await getList();
    const initIndexs = Array.from(Array(11).keys());
    const [a, b] = initialAB(initIndexs);
    setOptionAIdx(a);
    setOptionBIdx(b);
    let _list = removeFromList(initIndexs, a);
    _list = removeFromList(_list, b);
    setFoodList(_foodList);
    updateIndexList(_list);
  }, []);

  const onClickOptionA = () => {
    vote(foodList[optionAIdx].id, foodList[optionBIdx].id);
    if (indexList.length === 0) {
      setDone(true);
      return;
    }
    const newOptionB = chooseFromList(indexList);
    const _list = removeFromList(indexList, newOptionB);
    updateIndexList(_list);
    setOptionBIdx(newOptionB);
    setCountPlayed(countPlayed + 1);
  };

  const onClickOptionB = () => {
    vote(foodList[optionBIdx].id, foodList[optionAIdx].id);
    if (indexList.length === 0) {
      setDone(true);
      return;
    }
    const newOptionA = chooseFromList(indexList);
    const _list = removeFromList(indexList, newOptionA);
    updateIndexList(_list);
    setOptionAIdx(newOptionA);
    setCountPlayed(countPlayed + 1);
  };

  return (
    <>
      <div style={{ paddingTop: '6%', paddingBottom: '4%', textAlign: 'center', fontSize: 36 }}>
        哪個好吃
      </div>
      {foodList.length === 0 ? (
        <CenterMessage text='載入中...' />
      ) : done ? (
        <CenterMessage text='沒東西啦！' />
      ) : (
        <div style={{ textAlign: 'center' }}>
          <div style={{ padding: '3%', display: 'inline-block' }}>
            <Card style={{ maxWidth: '650' }}>
              <ButtonBase disabled={done} onClick={onClickOptionA}>
                <CardHeader>選項A</CardHeader>
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
                <CardHeader>選項B</CardHeader>
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
          {countPlayed > MIN_PLAYED_RESULT ? (
            <Button display='true' onClick={() => history.push('/ranking')}>
              {' '}
              看結果{' '}
            </Button>
          ) : (
            <></>
          )}
        </div>
      </div>
    </>
  );
}

export default Compare;
