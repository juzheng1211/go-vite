package trie

import (
	"github.com/syndtr/goleveldb/leveldb"
	"github.com/vitelabs/go-vite/chain_db/database"
	"github.com/vitelabs/go-vite/common/types"
	"github.com/vitelabs/go-vite/log15"
)

type Trie struct {
	db        *leveldb.DB
	cachePool *TrieNodePool
	log       log15.Logger

	RootHash *types.Hash
	Root     *TrieNode
}

func NewTrie(db *leveldb.DB, rootHash *types.Hash, pool *TrieNodePool) (*Trie, error) {
	trie := &Trie{
		db:       db,
		RootHash: rootHash,
		log:      log15.New("module", "vm_context"),
	}

	trie.loadFromDb()
	return trie, nil
}

func (trie *Trie) getNodeFromDb(key *types.Hash) *TrieNode {
	dbKey, _ := database.EncodeKey(database.DBKP_TRIE_NODE, key.Bytes())
	value, err := trie.db.Get(dbKey, nil)
	if err != nil {
		trie.log.Error("Query trie node failed from the database, error is "+err.Error(), "method", "getNodeFromDb")
		return nil
	}
	trieNode := &TrieNode{}
	dsErr := trieNode.DbDeserialize(value)
	if dsErr != nil {
		trie.log.Error("Deserialize trie node  failed, error is "+err.Error(), "method", "getNodeFromDb")
		return nil
	}

	return trieNode
}

func (trie *Trie) getNode(key *types.Hash) *TrieNode {
	node := trie.cachePool.Get(key)
	if node != nil {
		return node
	}

	node = trie.getNodeFromDb(key)
	if node != nil {
		trie.cachePool.Set(key, node)
	}
	return node
}

func (trie *Trie) loadFromDb() error {
	return nil
}

func (trie *Trie) computeHash() {

}

func (trie *Trie) Copy() *Trie {
	return &Trie{
		Root: trie.Root.Copy(),
	}
}

func (trie *Trie) Save() {

}

func (trie *Trie) SetValue(key []byte, value []byte) {

}

func (trie *Trie) GetValue(key []byte) []byte {
	return nil
}
