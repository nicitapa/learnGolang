1.SELECT model,speed,hd FROM PC
  WHERE price < 500

2.SELECTT maker FROM Product
  WHERE type ='Printer'
  GROUP BY maker

3.SELECT model,ram,sceen FROM Laptop
  WHERE price > 1000

4.SELECT * FROM Printer
  WHERE color = 'y'

5.SELECT model,speed,hd FROM PC
  WHERE cd in ('12x', '24x') AND price < 600

6.SELECT p.maker Maker, l.speed as speed FROM PRODUCT p
  left JOIN Lapyop l on p.model = l.model
  WHERE l.hd >= 10
  GROUP BY p.maker, l.speed

7.SELECT p.model, pc.price
FROM Product p
         JOIN PC pc ON p.model = pc.model
 WHERE p.maker = 'B'
 UNION
 SELECT p.model, l.price
 FROM Product p
         JOIN Laptop l ON p.model = l.model
WHERE p.maker = 'B'
UNION
SELECT p.model, pr.price
FROM Product p
         JOIN Printer pr ON p.model = pr.model
WHERE p.maker = 'B';

8.SELECT DISTINCT maker
FROM Product
WHERE type = 'PC'
  AND maker NOT IN (
    SELECT maker
    FROM Product
    WHERE type = 'Laptop'
);

9.SELECT DISTINCT p.maker
FROM Product p
         JOIN PC pc ON p.model = pc.model
WHERE pc.speed >= 450;

10.SELECT model, price
FROM Printer
WHERE price = (SELECT MAX(price) FROM Printer);