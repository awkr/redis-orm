// Code generated by go-bindata.
// sources:
// tpl/conf.elastic.gogo
// tpl/conf.mongo.gogo
// tpl/conf.mssql.gogo
// tpl/conf.mysql.gogo
// tpl/conf.orm.gogo
// tpl/conf.redis.gogo
// tpl/object.db.gogo
// tpl/object.db.query.gogo
// tpl/object.db.read.gogo
// tpl/object.db.write.gogo
// tpl/object.elastic.gogo
// tpl/object.functions.gogo
// tpl/object.gogo
// tpl/object.index.gogo
// tpl/object.mongo.gogo
// tpl/object.primary.key.gogo
// tpl/object.range.gogo
// tpl/object.redis.gogo
// tpl/object.redis.manager.gogo
// tpl/object.redis.pipeline.gogo
// tpl/object.redis.read.gogo
// tpl/object.redis.sync.gogo
// tpl/object.redis.write.gogo
// tpl/object.relation.gogo
// tpl/object.unqiue.gogo
// tpl/query.gogo
// tpl/relation.db.read.gogo
// tpl/relation.functions.gogo
// tpl/relation.geo.gogo
// tpl/relation.geo.sync.gogo
// tpl/relation.gogo
// tpl/relation.list.gogo
// tpl/relation.list.sync.gogo
// tpl/relation.manager.gogo
// tpl/relation.pair.gogo
// tpl/relation.pair.sync.gogo
// tpl/relation.pipeline.gogo
// tpl/relation.set.gogo
// tpl/relation.set.sync.gogo
// tpl/relation.zset.gogo
// tpl/relation.zset.sync.gogo
// tpl/script.mysql.sql
// tpl/view.gogo
// DO NOT EDIT!

package tpl

import (
	"fmt"
	"io/ioutil"
	"os"
	"path/filepath"
	"strings"
)

// bindataRead reads the given file from disk. It returns an error on failure.
func bindataRead(path, name string) ([]byte, error) {
	buf, err := ioutil.ReadFile(path)
	if err != nil {
		err = fmt.Errorf("Error reading asset %s at %s: %v", name, path, err)
	}
	return buf, err
}

type asset struct {
	bytes []byte
	info  os.FileInfo
}

// tplConfElasticGogo reads file data from disk. It returns an error on failure.
func tplConfElasticGogo() (*asset, error) {
	path := "/Users/blue/letgo/src/github.com/ezbuy/redis-orm/tpl/conf.elastic.gogo"
	name := "tpl/conf.elastic.gogo"
	bytes, err := bindataRead(path, name)
	if err != nil {
		return nil, err
	}

	fi, err := os.Stat(path)
	if err != nil {
		err = fmt.Errorf("Error reading asset info %s at %s: %v", name, path, err)
	}

	a := &asset{bytes: bytes, info: fi}
	return a, err
}

// tplConfMongoGogo reads file data from disk. It returns an error on failure.
func tplConfMongoGogo() (*asset, error) {
	path := "/Users/blue/letgo/src/github.com/ezbuy/redis-orm/tpl/conf.mongo.gogo"
	name := "tpl/conf.mongo.gogo"
	bytes, err := bindataRead(path, name)
	if err != nil {
		return nil, err
	}

	fi, err := os.Stat(path)
	if err != nil {
		err = fmt.Errorf("Error reading asset info %s at %s: %v", name, path, err)
	}

	a := &asset{bytes: bytes, info: fi}
	return a, err
}

// tplConfMssqlGogo reads file data from disk. It returns an error on failure.
func tplConfMssqlGogo() (*asset, error) {
	path := "/Users/blue/letgo/src/github.com/ezbuy/redis-orm/tpl/conf.mssql.gogo"
	name := "tpl/conf.mssql.gogo"
	bytes, err := bindataRead(path, name)
	if err != nil {
		return nil, err
	}

	fi, err := os.Stat(path)
	if err != nil {
		err = fmt.Errorf("Error reading asset info %s at %s: %v", name, path, err)
	}

	a := &asset{bytes: bytes, info: fi}
	return a, err
}

// tplConfMysqlGogo reads file data from disk. It returns an error on failure.
func tplConfMysqlGogo() (*asset, error) {
	path := "/Users/blue/letgo/src/github.com/ezbuy/redis-orm/tpl/conf.mysql.gogo"
	name := "tpl/conf.mysql.gogo"
	bytes, err := bindataRead(path, name)
	if err != nil {
		return nil, err
	}

	fi, err := os.Stat(path)
	if err != nil {
		err = fmt.Errorf("Error reading asset info %s at %s: %v", name, path, err)
	}

	a := &asset{bytes: bytes, info: fi}
	return a, err
}

// tplConfOrmGogo reads file data from disk. It returns an error on failure.
func tplConfOrmGogo() (*asset, error) {
	path := "/Users/blue/letgo/src/github.com/ezbuy/redis-orm/tpl/conf.orm.gogo"
	name := "tpl/conf.orm.gogo"
	bytes, err := bindataRead(path, name)
	if err != nil {
		return nil, err
	}

	fi, err := os.Stat(path)
	if err != nil {
		err = fmt.Errorf("Error reading asset info %s at %s: %v", name, path, err)
	}

	a := &asset{bytes: bytes, info: fi}
	return a, err
}

// tplConfRedisGogo reads file data from disk. It returns an error on failure.
func tplConfRedisGogo() (*asset, error) {
	path := "/Users/blue/letgo/src/github.com/ezbuy/redis-orm/tpl/conf.redis.gogo"
	name := "tpl/conf.redis.gogo"
	bytes, err := bindataRead(path, name)
	if err != nil {
		return nil, err
	}

	fi, err := os.Stat(path)
	if err != nil {
		err = fmt.Errorf("Error reading asset info %s at %s: %v", name, path, err)
	}

	a := &asset{bytes: bytes, info: fi}
	return a, err
}

// tplObjectDbGogo reads file data from disk. It returns an error on failure.
func tplObjectDbGogo() (*asset, error) {
	path := "/Users/blue/letgo/src/github.com/ezbuy/redis-orm/tpl/object.db.gogo"
	name := "tpl/object.db.gogo"
	bytes, err := bindataRead(path, name)
	if err != nil {
		return nil, err
	}

	fi, err := os.Stat(path)
	if err != nil {
		err = fmt.Errorf("Error reading asset info %s at %s: %v", name, path, err)
	}

	a := &asset{bytes: bytes, info: fi}
	return a, err
}

// tplObjectDbQueryGogo reads file data from disk. It returns an error on failure.
func tplObjectDbQueryGogo() (*asset, error) {
	path := "/Users/blue/letgo/src/github.com/ezbuy/redis-orm/tpl/object.db.query.gogo"
	name := "tpl/object.db.query.gogo"
	bytes, err := bindataRead(path, name)
	if err != nil {
		return nil, err
	}

	fi, err := os.Stat(path)
	if err != nil {
		err = fmt.Errorf("Error reading asset info %s at %s: %v", name, path, err)
	}

	a := &asset{bytes: bytes, info: fi}
	return a, err
}

// tplObjectDbReadGogo reads file data from disk. It returns an error on failure.
func tplObjectDbReadGogo() (*asset, error) {
	path := "/Users/blue/letgo/src/github.com/ezbuy/redis-orm/tpl/object.db.read.gogo"
	name := "tpl/object.db.read.gogo"
	bytes, err := bindataRead(path, name)
	if err != nil {
		return nil, err
	}

	fi, err := os.Stat(path)
	if err != nil {
		err = fmt.Errorf("Error reading asset info %s at %s: %v", name, path, err)
	}

	a := &asset{bytes: bytes, info: fi}
	return a, err
}

// tplObjectDbWriteGogo reads file data from disk. It returns an error on failure.
func tplObjectDbWriteGogo() (*asset, error) {
	path := "/Users/blue/letgo/src/github.com/ezbuy/redis-orm/tpl/object.db.write.gogo"
	name := "tpl/object.db.write.gogo"
	bytes, err := bindataRead(path, name)
	if err != nil {
		return nil, err
	}

	fi, err := os.Stat(path)
	if err != nil {
		err = fmt.Errorf("Error reading asset info %s at %s: %v", name, path, err)
	}

	a := &asset{bytes: bytes, info: fi}
	return a, err
}

// tplObjectElasticGogo reads file data from disk. It returns an error on failure.
func tplObjectElasticGogo() (*asset, error) {
	path := "/Users/blue/letgo/src/github.com/ezbuy/redis-orm/tpl/object.elastic.gogo"
	name := "tpl/object.elastic.gogo"
	bytes, err := bindataRead(path, name)
	if err != nil {
		return nil, err
	}

	fi, err := os.Stat(path)
	if err != nil {
		err = fmt.Errorf("Error reading asset info %s at %s: %v", name, path, err)
	}

	a := &asset{bytes: bytes, info: fi}
	return a, err
}

// tplObjectFunctionsGogo reads file data from disk. It returns an error on failure.
func tplObjectFunctionsGogo() (*asset, error) {
	path := "/Users/blue/letgo/src/github.com/ezbuy/redis-orm/tpl/object.functions.gogo"
	name := "tpl/object.functions.gogo"
	bytes, err := bindataRead(path, name)
	if err != nil {
		return nil, err
	}

	fi, err := os.Stat(path)
	if err != nil {
		err = fmt.Errorf("Error reading asset info %s at %s: %v", name, path, err)
	}

	a := &asset{bytes: bytes, info: fi}
	return a, err
}

// tplObjectGogo reads file data from disk. It returns an error on failure.
func tplObjectGogo() (*asset, error) {
	path := "/Users/blue/letgo/src/github.com/ezbuy/redis-orm/tpl/object.gogo"
	name := "tpl/object.gogo"
	bytes, err := bindataRead(path, name)
	if err != nil {
		return nil, err
	}

	fi, err := os.Stat(path)
	if err != nil {
		err = fmt.Errorf("Error reading asset info %s at %s: %v", name, path, err)
	}

	a := &asset{bytes: bytes, info: fi}
	return a, err
}

// tplObjectIndexGogo reads file data from disk. It returns an error on failure.
func tplObjectIndexGogo() (*asset, error) {
	path := "/Users/blue/letgo/src/github.com/ezbuy/redis-orm/tpl/object.index.gogo"
	name := "tpl/object.index.gogo"
	bytes, err := bindataRead(path, name)
	if err != nil {
		return nil, err
	}

	fi, err := os.Stat(path)
	if err != nil {
		err = fmt.Errorf("Error reading asset info %s at %s: %v", name, path, err)
	}

	a := &asset{bytes: bytes, info: fi}
	return a, err
}

// tplObjectMongoGogo reads file data from disk. It returns an error on failure.
func tplObjectMongoGogo() (*asset, error) {
	path := "/Users/blue/letgo/src/github.com/ezbuy/redis-orm/tpl/object.mongo.gogo"
	name := "tpl/object.mongo.gogo"
	bytes, err := bindataRead(path, name)
	if err != nil {
		return nil, err
	}

	fi, err := os.Stat(path)
	if err != nil {
		err = fmt.Errorf("Error reading asset info %s at %s: %v", name, path, err)
	}

	a := &asset{bytes: bytes, info: fi}
	return a, err
}

// tplObjectPrimaryKeyGogo reads file data from disk. It returns an error on failure.
func tplObjectPrimaryKeyGogo() (*asset, error) {
	path := "/Users/blue/letgo/src/github.com/ezbuy/redis-orm/tpl/object.primary.key.gogo"
	name := "tpl/object.primary.key.gogo"
	bytes, err := bindataRead(path, name)
	if err != nil {
		return nil, err
	}

	fi, err := os.Stat(path)
	if err != nil {
		err = fmt.Errorf("Error reading asset info %s at %s: %v", name, path, err)
	}

	a := &asset{bytes: bytes, info: fi}
	return a, err
}

// tplObjectRangeGogo reads file data from disk. It returns an error on failure.
func tplObjectRangeGogo() (*asset, error) {
	path := "/Users/blue/letgo/src/github.com/ezbuy/redis-orm/tpl/object.range.gogo"
	name := "tpl/object.range.gogo"
	bytes, err := bindataRead(path, name)
	if err != nil {
		return nil, err
	}

	fi, err := os.Stat(path)
	if err != nil {
		err = fmt.Errorf("Error reading asset info %s at %s: %v", name, path, err)
	}

	a := &asset{bytes: bytes, info: fi}
	return a, err
}

// tplObjectRedisGogo reads file data from disk. It returns an error on failure.
func tplObjectRedisGogo() (*asset, error) {
	path := "/Users/blue/letgo/src/github.com/ezbuy/redis-orm/tpl/object.redis.gogo"
	name := "tpl/object.redis.gogo"
	bytes, err := bindataRead(path, name)
	if err != nil {
		return nil, err
	}

	fi, err := os.Stat(path)
	if err != nil {
		err = fmt.Errorf("Error reading asset info %s at %s: %v", name, path, err)
	}

	a := &asset{bytes: bytes, info: fi}
	return a, err
}

// tplObjectRedisManagerGogo reads file data from disk. It returns an error on failure.
func tplObjectRedisManagerGogo() (*asset, error) {
	path := "/Users/blue/letgo/src/github.com/ezbuy/redis-orm/tpl/object.redis.manager.gogo"
	name := "tpl/object.redis.manager.gogo"
	bytes, err := bindataRead(path, name)
	if err != nil {
		return nil, err
	}

	fi, err := os.Stat(path)
	if err != nil {
		err = fmt.Errorf("Error reading asset info %s at %s: %v", name, path, err)
	}

	a := &asset{bytes: bytes, info: fi}
	return a, err
}

// tplObjectRedisPipelineGogo reads file data from disk. It returns an error on failure.
func tplObjectRedisPipelineGogo() (*asset, error) {
	path := "/Users/blue/letgo/src/github.com/ezbuy/redis-orm/tpl/object.redis.pipeline.gogo"
	name := "tpl/object.redis.pipeline.gogo"
	bytes, err := bindataRead(path, name)
	if err != nil {
		return nil, err
	}

	fi, err := os.Stat(path)
	if err != nil {
		err = fmt.Errorf("Error reading asset info %s at %s: %v", name, path, err)
	}

	a := &asset{bytes: bytes, info: fi}
	return a, err
}

// tplObjectRedisReadGogo reads file data from disk. It returns an error on failure.
func tplObjectRedisReadGogo() (*asset, error) {
	path := "/Users/blue/letgo/src/github.com/ezbuy/redis-orm/tpl/object.redis.read.gogo"
	name := "tpl/object.redis.read.gogo"
	bytes, err := bindataRead(path, name)
	if err != nil {
		return nil, err
	}

	fi, err := os.Stat(path)
	if err != nil {
		err = fmt.Errorf("Error reading asset info %s at %s: %v", name, path, err)
	}

	a := &asset{bytes: bytes, info: fi}
	return a, err
}

// tplObjectRedisSyncGogo reads file data from disk. It returns an error on failure.
func tplObjectRedisSyncGogo() (*asset, error) {
	path := "/Users/blue/letgo/src/github.com/ezbuy/redis-orm/tpl/object.redis.sync.gogo"
	name := "tpl/object.redis.sync.gogo"
	bytes, err := bindataRead(path, name)
	if err != nil {
		return nil, err
	}

	fi, err := os.Stat(path)
	if err != nil {
		err = fmt.Errorf("Error reading asset info %s at %s: %v", name, path, err)
	}

	a := &asset{bytes: bytes, info: fi}
	return a, err
}

// tplObjectRedisWriteGogo reads file data from disk. It returns an error on failure.
func tplObjectRedisWriteGogo() (*asset, error) {
	path := "/Users/blue/letgo/src/github.com/ezbuy/redis-orm/tpl/object.redis.write.gogo"
	name := "tpl/object.redis.write.gogo"
	bytes, err := bindataRead(path, name)
	if err != nil {
		return nil, err
	}

	fi, err := os.Stat(path)
	if err != nil {
		err = fmt.Errorf("Error reading asset info %s at %s: %v", name, path, err)
	}

	a := &asset{bytes: bytes, info: fi}
	return a, err
}

// tplObjectRelationGogo reads file data from disk. It returns an error on failure.
func tplObjectRelationGogo() (*asset, error) {
	path := "/Users/blue/letgo/src/github.com/ezbuy/redis-orm/tpl/object.relation.gogo"
	name := "tpl/object.relation.gogo"
	bytes, err := bindataRead(path, name)
	if err != nil {
		return nil, err
	}

	fi, err := os.Stat(path)
	if err != nil {
		err = fmt.Errorf("Error reading asset info %s at %s: %v", name, path, err)
	}

	a := &asset{bytes: bytes, info: fi}
	return a, err
}

// tplObjectUnqiueGogo reads file data from disk. It returns an error on failure.
func tplObjectUnqiueGogo() (*asset, error) {
	path := "/Users/blue/letgo/src/github.com/ezbuy/redis-orm/tpl/object.unqiue.gogo"
	name := "tpl/object.unqiue.gogo"
	bytes, err := bindataRead(path, name)
	if err != nil {
		return nil, err
	}

	fi, err := os.Stat(path)
	if err != nil {
		err = fmt.Errorf("Error reading asset info %s at %s: %v", name, path, err)
	}

	a := &asset{bytes: bytes, info: fi}
	return a, err
}

// tplQueryGogo reads file data from disk. It returns an error on failure.
func tplQueryGogo() (*asset, error) {
	path := "/Users/blue/letgo/src/github.com/ezbuy/redis-orm/tpl/query.gogo"
	name := "tpl/query.gogo"
	bytes, err := bindataRead(path, name)
	if err != nil {
		return nil, err
	}

	fi, err := os.Stat(path)
	if err != nil {
		err = fmt.Errorf("Error reading asset info %s at %s: %v", name, path, err)
	}

	a := &asset{bytes: bytes, info: fi}
	return a, err
}

// tplRelationDbReadGogo reads file data from disk. It returns an error on failure.
func tplRelationDbReadGogo() (*asset, error) {
	path := "/Users/blue/letgo/src/github.com/ezbuy/redis-orm/tpl/relation.db.read.gogo"
	name := "tpl/relation.db.read.gogo"
	bytes, err := bindataRead(path, name)
	if err != nil {
		return nil, err
	}

	fi, err := os.Stat(path)
	if err != nil {
		err = fmt.Errorf("Error reading asset info %s at %s: %v", name, path, err)
	}

	a := &asset{bytes: bytes, info: fi}
	return a, err
}

// tplRelationFunctionsGogo reads file data from disk. It returns an error on failure.
func tplRelationFunctionsGogo() (*asset, error) {
	path := "/Users/blue/letgo/src/github.com/ezbuy/redis-orm/tpl/relation.functions.gogo"
	name := "tpl/relation.functions.gogo"
	bytes, err := bindataRead(path, name)
	if err != nil {
		return nil, err
	}

	fi, err := os.Stat(path)
	if err != nil {
		err = fmt.Errorf("Error reading asset info %s at %s: %v", name, path, err)
	}

	a := &asset{bytes: bytes, info: fi}
	return a, err
}

// tplRelationGeoGogo reads file data from disk. It returns an error on failure.
func tplRelationGeoGogo() (*asset, error) {
	path := "/Users/blue/letgo/src/github.com/ezbuy/redis-orm/tpl/relation.geo.gogo"
	name := "tpl/relation.geo.gogo"
	bytes, err := bindataRead(path, name)
	if err != nil {
		return nil, err
	}

	fi, err := os.Stat(path)
	if err != nil {
		err = fmt.Errorf("Error reading asset info %s at %s: %v", name, path, err)
	}

	a := &asset{bytes: bytes, info: fi}
	return a, err
}

// tplRelationGeoSyncGogo reads file data from disk. It returns an error on failure.
func tplRelationGeoSyncGogo() (*asset, error) {
	path := "/Users/blue/letgo/src/github.com/ezbuy/redis-orm/tpl/relation.geo.sync.gogo"
	name := "tpl/relation.geo.sync.gogo"
	bytes, err := bindataRead(path, name)
	if err != nil {
		return nil, err
	}

	fi, err := os.Stat(path)
	if err != nil {
		err = fmt.Errorf("Error reading asset info %s at %s: %v", name, path, err)
	}

	a := &asset{bytes: bytes, info: fi}
	return a, err
}

// tplRelationGogo reads file data from disk. It returns an error on failure.
func tplRelationGogo() (*asset, error) {
	path := "/Users/blue/letgo/src/github.com/ezbuy/redis-orm/tpl/relation.gogo"
	name := "tpl/relation.gogo"
	bytes, err := bindataRead(path, name)
	if err != nil {
		return nil, err
	}

	fi, err := os.Stat(path)
	if err != nil {
		err = fmt.Errorf("Error reading asset info %s at %s: %v", name, path, err)
	}

	a := &asset{bytes: bytes, info: fi}
	return a, err
}

// tplRelationListGogo reads file data from disk. It returns an error on failure.
func tplRelationListGogo() (*asset, error) {
	path := "/Users/blue/letgo/src/github.com/ezbuy/redis-orm/tpl/relation.list.gogo"
	name := "tpl/relation.list.gogo"
	bytes, err := bindataRead(path, name)
	if err != nil {
		return nil, err
	}

	fi, err := os.Stat(path)
	if err != nil {
		err = fmt.Errorf("Error reading asset info %s at %s: %v", name, path, err)
	}

	a := &asset{bytes: bytes, info: fi}
	return a, err
}

// tplRelationListSyncGogo reads file data from disk. It returns an error on failure.
func tplRelationListSyncGogo() (*asset, error) {
	path := "/Users/blue/letgo/src/github.com/ezbuy/redis-orm/tpl/relation.list.sync.gogo"
	name := "tpl/relation.list.sync.gogo"
	bytes, err := bindataRead(path, name)
	if err != nil {
		return nil, err
	}

	fi, err := os.Stat(path)
	if err != nil {
		err = fmt.Errorf("Error reading asset info %s at %s: %v", name, path, err)
	}

	a := &asset{bytes: bytes, info: fi}
	return a, err
}

// tplRelationManagerGogo reads file data from disk. It returns an error on failure.
func tplRelationManagerGogo() (*asset, error) {
	path := "/Users/blue/letgo/src/github.com/ezbuy/redis-orm/tpl/relation.manager.gogo"
	name := "tpl/relation.manager.gogo"
	bytes, err := bindataRead(path, name)
	if err != nil {
		return nil, err
	}

	fi, err := os.Stat(path)
	if err != nil {
		err = fmt.Errorf("Error reading asset info %s at %s: %v", name, path, err)
	}

	a := &asset{bytes: bytes, info: fi}
	return a, err
}

// tplRelationPairGogo reads file data from disk. It returns an error on failure.
func tplRelationPairGogo() (*asset, error) {
	path := "/Users/blue/letgo/src/github.com/ezbuy/redis-orm/tpl/relation.pair.gogo"
	name := "tpl/relation.pair.gogo"
	bytes, err := bindataRead(path, name)
	if err != nil {
		return nil, err
	}

	fi, err := os.Stat(path)
	if err != nil {
		err = fmt.Errorf("Error reading asset info %s at %s: %v", name, path, err)
	}

	a := &asset{bytes: bytes, info: fi}
	return a, err
}

// tplRelationPairSyncGogo reads file data from disk. It returns an error on failure.
func tplRelationPairSyncGogo() (*asset, error) {
	path := "/Users/blue/letgo/src/github.com/ezbuy/redis-orm/tpl/relation.pair.sync.gogo"
	name := "tpl/relation.pair.sync.gogo"
	bytes, err := bindataRead(path, name)
	if err != nil {
		return nil, err
	}

	fi, err := os.Stat(path)
	if err != nil {
		err = fmt.Errorf("Error reading asset info %s at %s: %v", name, path, err)
	}

	a := &asset{bytes: bytes, info: fi}
	return a, err
}

// tplRelationPipelineGogo reads file data from disk. It returns an error on failure.
func tplRelationPipelineGogo() (*asset, error) {
	path := "/Users/blue/letgo/src/github.com/ezbuy/redis-orm/tpl/relation.pipeline.gogo"
	name := "tpl/relation.pipeline.gogo"
	bytes, err := bindataRead(path, name)
	if err != nil {
		return nil, err
	}

	fi, err := os.Stat(path)
	if err != nil {
		err = fmt.Errorf("Error reading asset info %s at %s: %v", name, path, err)
	}

	a := &asset{bytes: bytes, info: fi}
	return a, err
}

// tplRelationSetGogo reads file data from disk. It returns an error on failure.
func tplRelationSetGogo() (*asset, error) {
	path := "/Users/blue/letgo/src/github.com/ezbuy/redis-orm/tpl/relation.set.gogo"
	name := "tpl/relation.set.gogo"
	bytes, err := bindataRead(path, name)
	if err != nil {
		return nil, err
	}

	fi, err := os.Stat(path)
	if err != nil {
		err = fmt.Errorf("Error reading asset info %s at %s: %v", name, path, err)
	}

	a := &asset{bytes: bytes, info: fi}
	return a, err
}

// tplRelationSetSyncGogo reads file data from disk. It returns an error on failure.
func tplRelationSetSyncGogo() (*asset, error) {
	path := "/Users/blue/letgo/src/github.com/ezbuy/redis-orm/tpl/relation.set.sync.gogo"
	name := "tpl/relation.set.sync.gogo"
	bytes, err := bindataRead(path, name)
	if err != nil {
		return nil, err
	}

	fi, err := os.Stat(path)
	if err != nil {
		err = fmt.Errorf("Error reading asset info %s at %s: %v", name, path, err)
	}

	a := &asset{bytes: bytes, info: fi}
	return a, err
}

// tplRelationZsetGogo reads file data from disk. It returns an error on failure.
func tplRelationZsetGogo() (*asset, error) {
	path := "/Users/blue/letgo/src/github.com/ezbuy/redis-orm/tpl/relation.zset.gogo"
	name := "tpl/relation.zset.gogo"
	bytes, err := bindataRead(path, name)
	if err != nil {
		return nil, err
	}

	fi, err := os.Stat(path)
	if err != nil {
		err = fmt.Errorf("Error reading asset info %s at %s: %v", name, path, err)
	}

	a := &asset{bytes: bytes, info: fi}
	return a, err
}

// tplRelationZsetSyncGogo reads file data from disk. It returns an error on failure.
func tplRelationZsetSyncGogo() (*asset, error) {
	path := "/Users/blue/letgo/src/github.com/ezbuy/redis-orm/tpl/relation.zset.sync.gogo"
	name := "tpl/relation.zset.sync.gogo"
	bytes, err := bindataRead(path, name)
	if err != nil {
		return nil, err
	}

	fi, err := os.Stat(path)
	if err != nil {
		err = fmt.Errorf("Error reading asset info %s at %s: %v", name, path, err)
	}

	a := &asset{bytes: bytes, info: fi}
	return a, err
}

// tplScriptMysqlSql reads file data from disk. It returns an error on failure.
func tplScriptMysqlSql() (*asset, error) {
	path := "/Users/blue/letgo/src/github.com/ezbuy/redis-orm/tpl/script.mysql.sql"
	name := "tpl/script.mysql.sql"
	bytes, err := bindataRead(path, name)
	if err != nil {
		return nil, err
	}

	fi, err := os.Stat(path)
	if err != nil {
		err = fmt.Errorf("Error reading asset info %s at %s: %v", name, path, err)
	}

	a := &asset{bytes: bytes, info: fi}
	return a, err
}

// tplViewGogo reads file data from disk. It returns an error on failure.
func tplViewGogo() (*asset, error) {
	path := "/Users/blue/letgo/src/github.com/ezbuy/redis-orm/tpl/view.gogo"
	name := "tpl/view.gogo"
	bytes, err := bindataRead(path, name)
	if err != nil {
		return nil, err
	}

	fi, err := os.Stat(path)
	if err != nil {
		err = fmt.Errorf("Error reading asset info %s at %s: %v", name, path, err)
	}

	a := &asset{bytes: bytes, info: fi}
	return a, err
}

// Asset loads and returns the asset for the given name.
// It returns an error if the asset could not be found or
// could not be loaded.
func Asset(name string) ([]byte, error) {
	cannonicalName := strings.Replace(name, "\\", "/", -1)
	if f, ok := _bindata[cannonicalName]; ok {
		a, err := f()
		if err != nil {
			return nil, fmt.Errorf("Asset %s can't read by error: %v", name, err)
		}
		return a.bytes, nil
	}
	return nil, fmt.Errorf("Asset %s not found", name)
}

// MustAsset is like Asset but panics when Asset would return an error.
// It simplifies safe initialization of global variables.
func MustAsset(name string) []byte {
	a, err := Asset(name)
	if err != nil {
		panic("asset: Asset(" + name + "): " + err.Error())
	}

	return a
}

// AssetInfo loads and returns the asset info for the given name.
// It returns an error if the asset could not be found or
// could not be loaded.
func AssetInfo(name string) (os.FileInfo, error) {
	cannonicalName := strings.Replace(name, "\\", "/", -1)
	if f, ok := _bindata[cannonicalName]; ok {
		a, err := f()
		if err != nil {
			return nil, fmt.Errorf("AssetInfo %s can't read by error: %v", name, err)
		}
		return a.info, nil
	}
	return nil, fmt.Errorf("AssetInfo %s not found", name)
}

// AssetNames returns the names of the assets.
func AssetNames() []string {
	names := make([]string, 0, len(_bindata))
	for name := range _bindata {
		names = append(names, name)
	}
	return names
}

// _bindata is a table, holding each asset generator, mapped to its name.
var _bindata = map[string]func() (*asset, error){
	"tpl/conf.elastic.gogo": tplConfElasticGogo,
	"tpl/conf.mongo.gogo": tplConfMongoGogo,
	"tpl/conf.mssql.gogo": tplConfMssqlGogo,
	"tpl/conf.mysql.gogo": tplConfMysqlGogo,
	"tpl/conf.orm.gogo": tplConfOrmGogo,
	"tpl/conf.redis.gogo": tplConfRedisGogo,
	"tpl/object.db.gogo": tplObjectDbGogo,
	"tpl/object.db.query.gogo": tplObjectDbQueryGogo,
	"tpl/object.db.read.gogo": tplObjectDbReadGogo,
	"tpl/object.db.write.gogo": tplObjectDbWriteGogo,
	"tpl/object.elastic.gogo": tplObjectElasticGogo,
	"tpl/object.functions.gogo": tplObjectFunctionsGogo,
	"tpl/object.gogo": tplObjectGogo,
	"tpl/object.index.gogo": tplObjectIndexGogo,
	"tpl/object.mongo.gogo": tplObjectMongoGogo,
	"tpl/object.primary.key.gogo": tplObjectPrimaryKeyGogo,
	"tpl/object.range.gogo": tplObjectRangeGogo,
	"tpl/object.redis.gogo": tplObjectRedisGogo,
	"tpl/object.redis.manager.gogo": tplObjectRedisManagerGogo,
	"tpl/object.redis.pipeline.gogo": tplObjectRedisPipelineGogo,
	"tpl/object.redis.read.gogo": tplObjectRedisReadGogo,
	"tpl/object.redis.sync.gogo": tplObjectRedisSyncGogo,
	"tpl/object.redis.write.gogo": tplObjectRedisWriteGogo,
	"tpl/object.relation.gogo": tplObjectRelationGogo,
	"tpl/object.unqiue.gogo": tplObjectUnqiueGogo,
	"tpl/query.gogo": tplQueryGogo,
	"tpl/relation.db.read.gogo": tplRelationDbReadGogo,
	"tpl/relation.functions.gogo": tplRelationFunctionsGogo,
	"tpl/relation.geo.gogo": tplRelationGeoGogo,
	"tpl/relation.geo.sync.gogo": tplRelationGeoSyncGogo,
	"tpl/relation.gogo": tplRelationGogo,
	"tpl/relation.list.gogo": tplRelationListGogo,
	"tpl/relation.list.sync.gogo": tplRelationListSyncGogo,
	"tpl/relation.manager.gogo": tplRelationManagerGogo,
	"tpl/relation.pair.gogo": tplRelationPairGogo,
	"tpl/relation.pair.sync.gogo": tplRelationPairSyncGogo,
	"tpl/relation.pipeline.gogo": tplRelationPipelineGogo,
	"tpl/relation.set.gogo": tplRelationSetGogo,
	"tpl/relation.set.sync.gogo": tplRelationSetSyncGogo,
	"tpl/relation.zset.gogo": tplRelationZsetGogo,
	"tpl/relation.zset.sync.gogo": tplRelationZsetSyncGogo,
	"tpl/script.mysql.sql": tplScriptMysqlSql,
	"tpl/view.gogo": tplViewGogo,
}

// AssetDir returns the file names below a certain
// directory embedded in the file by go-bindata.
// For example if you run go-bindata on data/... and data contains the
// following hierarchy:
//     data/
//       foo.txt
//       img/
//         a.png
//         b.png
// then AssetDir("data") would return []string{"foo.txt", "img"}
// AssetDir("data/img") would return []string{"a.png", "b.png"}
// AssetDir("foo.txt") and AssetDir("notexist") would return an error
// AssetDir("") will return []string{"data"}.
func AssetDir(name string) ([]string, error) {
	node := _bintree
	if len(name) != 0 {
		cannonicalName := strings.Replace(name, "\\", "/", -1)
		pathList := strings.Split(cannonicalName, "/")
		for _, p := range pathList {
			node = node.Children[p]
			if node == nil {
				return nil, fmt.Errorf("Asset %s not found", name)
			}
		}
	}
	if node.Func != nil {
		return nil, fmt.Errorf("Asset %s not found", name)
	}
	rv := make([]string, 0, len(node.Children))
	for childName := range node.Children {
		rv = append(rv, childName)
	}
	return rv, nil
}

type bintree struct {
	Func     func() (*asset, error)
	Children map[string]*bintree
}
var _bintree = &bintree{nil, map[string]*bintree{
	"tpl": &bintree{nil, map[string]*bintree{
		"conf.elastic.gogo": &bintree{tplConfElasticGogo, map[string]*bintree{}},
		"conf.mongo.gogo": &bintree{tplConfMongoGogo, map[string]*bintree{}},
		"conf.mssql.gogo": &bintree{tplConfMssqlGogo, map[string]*bintree{}},
		"conf.mysql.gogo": &bintree{tplConfMysqlGogo, map[string]*bintree{}},
		"conf.orm.gogo": &bintree{tplConfOrmGogo, map[string]*bintree{}},
		"conf.redis.gogo": &bintree{tplConfRedisGogo, map[string]*bintree{}},
		"object.db.gogo": &bintree{tplObjectDbGogo, map[string]*bintree{}},
		"object.db.query.gogo": &bintree{tplObjectDbQueryGogo, map[string]*bintree{}},
		"object.db.read.gogo": &bintree{tplObjectDbReadGogo, map[string]*bintree{}},
		"object.db.write.gogo": &bintree{tplObjectDbWriteGogo, map[string]*bintree{}},
		"object.elastic.gogo": &bintree{tplObjectElasticGogo, map[string]*bintree{}},
		"object.functions.gogo": &bintree{tplObjectFunctionsGogo, map[string]*bintree{}},
		"object.gogo": &bintree{tplObjectGogo, map[string]*bintree{}},
		"object.index.gogo": &bintree{tplObjectIndexGogo, map[string]*bintree{}},
		"object.mongo.gogo": &bintree{tplObjectMongoGogo, map[string]*bintree{}},
		"object.primary.key.gogo": &bintree{tplObjectPrimaryKeyGogo, map[string]*bintree{}},
		"object.range.gogo": &bintree{tplObjectRangeGogo, map[string]*bintree{}},
		"object.redis.gogo": &bintree{tplObjectRedisGogo, map[string]*bintree{}},
		"object.redis.manager.gogo": &bintree{tplObjectRedisManagerGogo, map[string]*bintree{}},
		"object.redis.pipeline.gogo": &bintree{tplObjectRedisPipelineGogo, map[string]*bintree{}},
		"object.redis.read.gogo": &bintree{tplObjectRedisReadGogo, map[string]*bintree{}},
		"object.redis.sync.gogo": &bintree{tplObjectRedisSyncGogo, map[string]*bintree{}},
		"object.redis.write.gogo": &bintree{tplObjectRedisWriteGogo, map[string]*bintree{}},
		"object.relation.gogo": &bintree{tplObjectRelationGogo, map[string]*bintree{}},
		"object.unqiue.gogo": &bintree{tplObjectUnqiueGogo, map[string]*bintree{}},
		"query.gogo": &bintree{tplQueryGogo, map[string]*bintree{}},
		"relation.db.read.gogo": &bintree{tplRelationDbReadGogo, map[string]*bintree{}},
		"relation.functions.gogo": &bintree{tplRelationFunctionsGogo, map[string]*bintree{}},
		"relation.geo.gogo": &bintree{tplRelationGeoGogo, map[string]*bintree{}},
		"relation.geo.sync.gogo": &bintree{tplRelationGeoSyncGogo, map[string]*bintree{}},
		"relation.gogo": &bintree{tplRelationGogo, map[string]*bintree{}},
		"relation.list.gogo": &bintree{tplRelationListGogo, map[string]*bintree{}},
		"relation.list.sync.gogo": &bintree{tplRelationListSyncGogo, map[string]*bintree{}},
		"relation.manager.gogo": &bintree{tplRelationManagerGogo, map[string]*bintree{}},
		"relation.pair.gogo": &bintree{tplRelationPairGogo, map[string]*bintree{}},
		"relation.pair.sync.gogo": &bintree{tplRelationPairSyncGogo, map[string]*bintree{}},
		"relation.pipeline.gogo": &bintree{tplRelationPipelineGogo, map[string]*bintree{}},
		"relation.set.gogo": &bintree{tplRelationSetGogo, map[string]*bintree{}},
		"relation.set.sync.gogo": &bintree{tplRelationSetSyncGogo, map[string]*bintree{}},
		"relation.zset.gogo": &bintree{tplRelationZsetGogo, map[string]*bintree{}},
		"relation.zset.sync.gogo": &bintree{tplRelationZsetSyncGogo, map[string]*bintree{}},
		"script.mysql.sql": &bintree{tplScriptMysqlSql, map[string]*bintree{}},
		"view.gogo": &bintree{tplViewGogo, map[string]*bintree{}},
	}},
}}

// RestoreAsset restores an asset under the given directory
func RestoreAsset(dir, name string) error {
	data, err := Asset(name)
	if err != nil {
		return err
	}
	info, err := AssetInfo(name)
	if err != nil {
		return err
	}
	err = os.MkdirAll(_filePath(dir, filepath.Dir(name)), os.FileMode(0755))
	if err != nil {
		return err
	}
	err = ioutil.WriteFile(_filePath(dir, name), data, info.Mode())
	if err != nil {
		return err
	}
	err = os.Chtimes(_filePath(dir, name), info.ModTime(), info.ModTime())
	if err != nil {
		return err
	}
	return nil
}

// RestoreAssets restores an asset under the given directory recursively
func RestoreAssets(dir, name string) error {
	children, err := AssetDir(name)
	// File
	if err != nil {
		return RestoreAsset(dir, name)
	}
	// Dir
	for _, child := range children {
		err = RestoreAssets(dir, filepath.Join(name, child))
		if err != nil {
			return err
		}
	}
	return nil
}

func _filePath(dir, name string) string {
	cannonicalName := strings.Replace(name, "\\", "/", -1)
	return filepath.Join(append([]string{dir}, strings.Split(cannonicalName, "/")...)...)
}

