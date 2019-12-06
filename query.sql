
$ go get github.com/labstack/echo
$ go get github.com/labstack/echo/middleware
$ go get github.com/go-sql-driver/mysql


-- Table structure for table `employee`
--

CREATE TABLE IF NOT EXISTS `person` (
  `id` int(11) NOT NULL AUTO_INCREMENT COMMENT 'person key',
  `name` varchar(255) NOT NULL COMMENT 'person name',
  `salary` double NOT NULL COMMENT 'person salary',
  `age` int(11) NOT NULL COMMENT 'person age',
  PRIMARY KEY (`id`)
) ENGINE=InnoDB  DEFAULT CHARSET=latin1 COMMENT='datatable demo table' 
AUTO_INCREMENT=1 ;